package file_service

import (
	"transactions_reader_stori/services/file_service/file_service_content_reader"
	"transactions_reader_stori/services/file_service/file_service_summary_generator"
)

type FileServiceFactoryI interface {
	NewFileService(
		fileContentReaderUseCase file_service_content_reader.FileContentReaderUseCaseI,
		fileSummaryGeneratorUseCase file_service_summary_generator.FileSummaryGeneratorUseCaseI) FileServiceI
}

type FileServiceFactory struct {
}

// NewFileService creates a new instance of FileService
func (factory *FileServiceFactory) NewFileService(
	fileContentReaderUseCase file_service_content_reader.FileContentReaderUseCaseI,
	fileSummaryGeneratorUseCase file_service_summary_generator.FileSummaryGeneratorUseCaseI) FileServiceI {
	return &FileService{
		fileContentReaderUseCase:    fileContentReaderUseCase,
		fileSummaryGeneratorUseCase: fileSummaryGeneratorUseCase,
	}
}
