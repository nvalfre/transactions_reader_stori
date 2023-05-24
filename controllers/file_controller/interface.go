package file_controller

import (
	"github.com/gin-gonic/gin"
)

type FileServiceI interface {
	ProcessFile(c *gin.Context) error
}
