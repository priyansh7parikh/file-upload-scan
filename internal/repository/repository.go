package repository

import (
	"sync"

	"github.com/priyansh7parikh/file-upload-scan/internal/model"

	"github.com/google/uuid"
)

type FileRepository struct {
	mu    sync.Mutex
	store map[uuid.UUID]model.File
}

func NewFileRepository() *FileRepository {
	return &FileRepository{
		store: make(map[uuid.UUID]model.File),
	}
}

func (r *FileRepository) Save(file model.File) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[file.ID] = file
}
