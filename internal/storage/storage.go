package storage

import (
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type TempStorage struct {
	BasePath string
}

func (s *TempStorage) Save(
	reader io.Reader,
) (string, int64, error) {

	fileName := uuid.New().String()
	fullPath := filepath.Join(s.BasePath, fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	size, err := io.Copy(file, reader)
	if err != nil {
		return "", 0, err
	}

	return fullPath, size, nil
}
