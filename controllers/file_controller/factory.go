package file_controller

import (
	"transactions_reader_stori/services/file_service"
)

type FileControllerFactoryI interface {
	NewFileController() *FileController
}

type FileControllerFactory struct {
	FileService file_service.FileServiceI
}

// NewFileController creates a new instance of FileController
func (fileControllerFactory *FileControllerFactory) NewFileController() *FileController {
	return &FileController{
		fileService: fileControllerFactory.FileService,
	}
}
