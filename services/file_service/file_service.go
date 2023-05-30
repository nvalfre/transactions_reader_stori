package file_service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"transactions_reader_stori/domain"
)

// ProcessFile processes the uploaded file
func (fs *FileService) ProcessFile(c *gin.Context) (*domain.Response, error) {
	file, err := c.FormFile("file")

	if err != nil {
		log.Println("Error occurred:", err)
		return nil, err
	}

	fileContent, err := fs.fileContentReaderUseCase.GetFileContentFromRequest(file)
	if err != nil {
		return &domain.Response{
			Status: http.StatusBadRequest,
			Body:   gin.H{"message": "Failed processing file content", "error": err.Error()},
		}, err
	}

	summary, err := fs.fileSummaryGeneratorUseCase.Execute(c, fileContent)
	if err != nil {
		return &domain.Response{
			Status: http.StatusBadRequest,
			Body:   gin.H{"message": "Failed executing summary generation", "error": err.Error()},
		}, err
	}

	return &domain.Response{
		Status: http.StatusOK,
		Body:   summary,
	}, nil
}
