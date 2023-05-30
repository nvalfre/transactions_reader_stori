package file_service_content_reader

import "mime/multipart"

type FileContentReaderUseCaseI interface {
	GetFileContent() ([]byte, error)
	GetFileContentFromRequest(file *multipart.FileHeader) ([]byte, error)
}
