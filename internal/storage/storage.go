package storage

import (
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/priyansh7parikh/file-upload-scan/internal/logger"
	"go.uber.org/zap"
)

type TempStorage struct {
	BasePath string
}

func (s *TempStorage) Save(reader io.Reader) (string, int64, error) {
	if err := os.MkdirAll(s.BasePath, 0755); err != nil {
		logger.Log.Error("failed to create temp directory",
			zap.String("path", s.BasePath),
			zap.Error(err),
		)
		return "", 0, err
	}

	fileName := uuid.New().String()
	fullPath := filepath.Join(s.BasePath, fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		logger.Log.Error("failed to create temp file",
			zap.String("path", fullPath),
			zap.Error(err),
		)
		return "", 0, err
	}
	defer file.Close()

	size, err := io.Copy(file, reader)
	if err != nil {
		logger.Log.Error("failed to write file",
			zap.String("path", fullPath),
			zap.Error(err),
		)
		return "", 0, err
	}

	logger.Log.Info("file written to temp storage",
		zap.String("path", fullPath),
		zap.Int64("size", size),
	)

	return fullPath, size, nil
}
