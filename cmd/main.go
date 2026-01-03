// @title File Upload Scan API
// @version 1.0
// @description API for file upload scanning service
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/priyansh7parikh/file-upload-scan/docs"
	"github.com/priyansh7parikh/file-upload-scan/internal/controller"
	"github.com/priyansh7parikh/file-upload-scan/internal/queue"
	"github.com/priyansh7parikh/file-upload-scan/internal/repository"
	service "github.com/priyansh7parikh/file-upload-scan/internal/services"
	"github.com/priyansh7parikh/file-upload-scan/internal/storage"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// ---- dependencies ----
	repo := repository.NewFileRepository()
	jobQueue := &queue.InMemoryQueue{}

	tempStorage := &storage.TempStorage{
		BasePath: "/tmp/uploads",
	}

	validator := &service.ValidationService{}
	uploadService := service.NewUploadService(
		validator,
		tempStorage,
		repo,
		jobQueue,
	)

	uploadHandler := controller.NewUploadHandler(uploadService)

	// ---- router ----
	mux := http.NewServeMux()
	mux.Handle("/upload", uploadHandler)

	// swagger
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// ---- graceful shutdown ----
	go func() {
		log.Println("🚀 Upload service running on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("⏳ shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown failed: %v", err)
	}

	log.Println("✅ server stopped gracefully")
}
