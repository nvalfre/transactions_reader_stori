package file_service

import (
	"transactions_reader_stori/services/file_service/file_service_content_reader"
	"transactions_reader_stori/services/file_service/file_service_summary_generator"
)

// FileService handles file processing HTTP endpoint
type FileService struct {
	fileContentReaderUseCase    file_service_content_reader.FileContentReaderUseCaseI
	fileSummaryGeneratorUseCase file_service_summary_generator.FileSummaryGeneratorUseCaseI
}
