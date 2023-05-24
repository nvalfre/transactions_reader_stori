package file_controller

import (
	"transactions_reader_stori/services/file_service"
)

// FileController handles file processing HTTP endpoint
type FileController struct {
	fileService file_service.FileServiceI
}
