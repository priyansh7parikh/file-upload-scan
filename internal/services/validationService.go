package service

import "mime/multipart"

const MaxUploadSize = 50 << 20 // 50 MB

var AllowedMimeTypes = map[string]bool{
	"image/png":       true,
	"image/jpeg":      true,
	"application/pdf": true,
}

type ValidationService struct{}

func (v *ValidationService) Validate(
	header *multipart.FileHeader,
) error {

	if header.Size > MaxUploadSize {
		return ErrFileTooLarge
	}

	contentType := header.Header.Get("Content-Type")
	if !AllowedMimeTypes[contentType] {
		return ErrInvalidFileType
	}

	return nil
}
