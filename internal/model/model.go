package model

import (
	"time"

	"github.com/google/uuid"
)

type FileStatus string

const (
	Uploaded FileStatus = "UPLOADED"
	Scanning FileStatus = "SCANNING"
	Clean    FileStatus = "CLEAN"
	Infected FileStatus = "INFECTED"
	Failed   FileStatus = "FAILED"
)

type File struct {
	ID           uuid.UUID
	OriginalName string
	MimeType     string
	Size         int64
	Status       FileStatus
	TempPath     string
	FinalPath    *string
	CreatedAt    time.Time
	ScannedAt    *time.Time
}

type ScanJob struct {
	ID         uuid.UUID
	FileID     uuid.UUID
	Status     string
	RetryCount int
	LastError  *string
	CreatedAt  time.Time
}
