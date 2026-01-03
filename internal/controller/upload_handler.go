package controller

import (
	"net/http"
	"time"

	"github.com/priyansh7parikh/file-upload-scan/internal/logger"
	service "github.com/priyansh7parikh/file-upload-scan/internal/services"
	"go.uber.org/zap"
)

type UploadHandler struct {
	service *service.UploadService
}

func NewUploadHandler(s *service.UploadService) *UploadHandler {
	return &UploadHandler{s}
}

// UploadFile godoc
// @Summary Upload a file securely
// @Description Uploads a file and schedules it for virus scanning
// @Tags files
// @Accept multipart/form-data
// @Produce plain
// @Param file formData file true "File to upload"
// @Success 202 {string} string "file_id"
// @Failure 400 {string} string "invalid request"
// @Failure 413 {string} string "file too large"
// @Failure 500 {string} string "internal error"
// @Router /upload [post]
func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	file, header, err := r.FormFile("file")
	if err != nil {
		logger.Log.Error("failed to read multipart file",
			zap.Error(err),
			zap.String("path", r.URL.Path),
		)
		http.Error(w, "invalid multipart request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	id, err := h.service.HandleUpload(r.Context(), file, header)
	if err != nil {
		logger.Log.Error("upload failed",
			zap.Error(err),
			zap.String("filename", header.Filename),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Log.Info("file uploaded successfully",
		zap.String("file_id", id.String()),
		zap.String("filename", header.Filename),
		zap.Duration("latency", time.Since(start)),
	)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(id.String()))
}
