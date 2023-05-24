package file_service

import (
	"github.com/gin-gonic/gin"
	"transactions_reader_stori/domain"
)

type FileServiceI interface {
	ProcessFile(c *gin.Context) (*domain.Response, error)
}
