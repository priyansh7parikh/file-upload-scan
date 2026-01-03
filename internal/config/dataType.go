package config

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
