package file_service_content_reader

import (
	"github.com/gin-gonic/gin"
)

type FileContentReaderUseCaseI interface {
	GetFileContent(c *gin.Context) ([]byte, error)
}
