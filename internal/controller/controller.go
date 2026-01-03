package controller

import (
	"errors"
	"net/http"

	service "github.com/priyansh7parikh/file-upload-scan/internal/services"
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
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "invalid multipart request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	id, err := h.service.HandleUpload(r.Context(), file, header)
	if err != nil {
		status := http.StatusInternalServerError

		switch {
		case errors.Is(err, service.ErrFileTooLarge):
			status = http.StatusRequestEntityTooLarge
		case errors.Is(err, service.ErrInvalidFileType):
			status = http.StatusBadRequest
		}

		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(id.String()))
}
