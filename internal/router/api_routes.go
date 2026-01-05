package router

import (
	"net/http"

	"github.com/priyansh7parikh/file-upload-scan/internal/auth"
	"github.com/priyansh7parikh/file-upload-scan/internal/controller"
	"github.com/priyansh7parikh/file-upload-scan/internal/queue"
	"github.com/priyansh7parikh/file-upload-scan/internal/repository"
	service "github.com/priyansh7parikh/file-upload-scan/internal/services"
	"github.com/priyansh7parikh/file-upload-scan/internal/storage"
)

func registerAPIRoutes(mux *http.ServeMux) {

	// -------- Infrastructure --------
	fileRepo := repository.NewFileRepository()
	jobQueue := &queue.InMemoryQueue{}

	tempStorage := &storage.TempStorage{
		BasePath: "/tmp/uploads",
	}

	validator := &service.ValidationService{}

	uploadService := service.NewUploadService(
		validator,
		tempStorage,
		fileRepo,
		jobQueue,
	)

	uploadHandler := controller.NewUploadHandler(uploadService)

	// -------- API v1 --------
	mux.Handle(
		"/api/v1/upload",
		auth.Middleware("user")(uploadHandler),
	)

	// -------- Health --------
	mux.HandleFunc("/health", healthHandler)
}
