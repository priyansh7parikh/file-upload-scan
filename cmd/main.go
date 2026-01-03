package main

import (
	"log"
	"net/http"

	"github.com/priyansh7parikh/file-upload-scan/internal/controller"
	"github.com/priyansh7parikh/file-upload-scan/internal/queue"
	"github.com/priyansh7parikh/file-upload-scan/internal/repository"
	service "github.com/priyansh7parikh/file-upload-scan/internal/services"
	"github.com/priyansh7parikh/file-upload-scan/internal/storage"
)

func main() {
	repo := repository.NewFileRepository()
	queue := &queue.InMemoryQueue{}
	storage := &storage.TempStorage{BasePath: "/tmp/uploads"}
	validator := &service.ValidationService{}

	uploadService := service.NewUploadService(
		validator, storage, repo, queue,
	)

	handler := controller.NewUploadHandler(uploadService)

	http.Handle("/upload", handler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
