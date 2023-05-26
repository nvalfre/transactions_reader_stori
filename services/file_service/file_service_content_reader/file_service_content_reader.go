package file_service_content_reader

import (
	"github.com/gin-gonic/gin"
	"log"
)

type FileContentReaderUseCase struct {
}

func (fs *FileContentReaderUseCase) GetFileContent(c *gin.Context) ([]byte, error) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error occurred:", err)
		return nil, err
	}

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
