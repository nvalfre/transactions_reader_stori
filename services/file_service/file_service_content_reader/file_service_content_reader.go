package file_service_content_reader

import (
	"log"
	"mime/multipart"
	"os"
)

type FileContentReaderUseCase struct {
}

func (fs *FileContentReaderUseCase) GetFileContent() ([]byte, error) {
	filePath := "balance.csv"

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error occurred:", err)
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("Error occurred:", err)
		return nil, err
	}

	fileSize := fileInfo.Size()
	fileContent := make([]byte, fileSize)

	_, err = file.Read(fileContent)
	if err != nil {
		log.Println("Error occurred:", err)
		return nil, err
	}

	return fileContent, nil
}

func (fs *FileContentReaderUseCase) GetFileContentFromRequest(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		log.Println("Error occurred:", err)
		return nil, err
	}
	defer src.Close()

	fileContent := make([]byte, file.Size)

	_, err = src.Read(fileContent)
	if err != nil {
		log.Println("Error occurred:", err)
		return nil, err
	}

	return fileContent, nil
}
