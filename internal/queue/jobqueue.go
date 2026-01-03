package queue

import "github.com/google/uuid"

type JobQueue interface {
	Enqueue(fileID uuid.UUID) error
}

type InMemoryQueue struct{}

func (q *InMemoryQueue) Enqueue(fileID uuid.UUID) error {
	// later replace with Redis
	return nil
}
