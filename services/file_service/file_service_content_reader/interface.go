package file_service_content_reader

type FileContentReaderUseCaseI interface {
	GetFileContent() ([]byte, error)
}
