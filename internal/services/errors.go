package service

import "errors"

var (
	ErrFileTooLarge    = errors.New("file size exceeds allowed limit")
	ErrInvalidFileType = errors.New("invalid file type")
	ErrUploadFailed    = errors.New("file upload failed")
)
