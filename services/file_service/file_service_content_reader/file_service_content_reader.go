package file_service_content_reader

import (
	"log"
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
