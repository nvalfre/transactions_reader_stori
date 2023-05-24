package file_controller

import (
	"transactions_reader_stori/services/file_service"
)

// NewFileController creates a new instance of FileController
func NewFileController(fileService file_service.FileServiceI) *FileController {
	return &FileController{
		fileService: fileService,
	}
}
