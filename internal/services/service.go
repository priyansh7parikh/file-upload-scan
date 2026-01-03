package service

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/priyansh7parikh/file-upload-scan/internal/model"
	"github.com/priyansh7parikh/file-upload-scan/internal/queue"
	"github.com/priyansh7parikh/file-upload-scan/internal/repository"
	"github.com/priyansh7parikh/file-upload-scan/internal/storage"
)

type UploadService struct {
	validator *ValidationService
	storage   *storage.TempStorage
	repo      *repository.FileRepository
	queue     queue.JobQueue
}

func NewUploadService(
	v *ValidationService,
	s *storage.TempStorage,
	r *repository.FileRepository,
	q queue.JobQueue,
) *UploadService {
	return &UploadService{v, s, r, q}
}

func (s *UploadService) HandleUpload(
	ctx context.Context,
	reader io.Reader,
	header *multipart.FileHeader,
) (uuid.UUID, error) {

	if err := s.validator.Validate(header); err != nil {
		return uuid.Nil, err
	}

	tempPath, size, err := s.storage.Save(reader)
	if err != nil {
		log.Println("error is %v", err)
		return uuid.Nil, err
	}

	file := model.File{
		ID:           uuid.New(),
		OriginalName: header.Filename,
		MimeType:     header.Header.Get("Content-Type"),
		Size:         size,
		Status:       model.Uploaded,
		TempPath:     tempPath,
		CreatedAt:    time.Now(),
	}

	s.repo.Save(file)
	s.queue.Enqueue(file.ID)

	return file.ID, nil
}
