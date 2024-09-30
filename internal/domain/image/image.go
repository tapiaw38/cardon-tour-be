package domain

import "mime/multipart"

type (
	ImageFile struct {
		File       multipart.File
		FileHeader *multipart.FileHeader
	}
)
