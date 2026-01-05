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

	"github.com/priyansh7parikh/file-upload-scan/internal/logger"
	"github.com/priyansh7parikh/file-upload-scan/internal/router"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	logger.Init(env)
	defer logger.Log.Sync()

	// ---- dependencies ----
	// repo := repository.NewFileRepository()
	// jobQueue := &queue.InMemoryQueue{}

	// tempStorage := &storage.TempStorage{
	// 	BasePath: "/tmp/uploads",
	// }

	// validator := &service.ValidationService{}
	// uploadService := service.NewUploadService(
	// 	validator,
	// 	tempStorage,
	// 	repo,
	// 	jobQueue,
	// )

	// uploadHandler := controller.NewUploadHandler(uploadService)

	// ---- router ----
	handler := router.New()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// ---- graceful shutdown ----
	go func() {
		log.Println("Upload service running on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown failed: %v", err)
	}

	log.Println("server stopped gracefully")
}
