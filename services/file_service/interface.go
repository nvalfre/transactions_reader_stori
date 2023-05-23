package file_service

import (
	"github.com/gin-gonic/gin"
)

type FileServiceI interface {
	ProcessFile(c *gin.Context)
}
