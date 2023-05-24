package file_controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

// ProcessFile processes the uploaded file
func (h *FileController) ProcessFile(c *gin.Context) {
	res, err := h.fileService.ProcessFile(c)
	if err != nil {
		log.Println("Error occurred:", err)
	}
	c.JSON(res.Status, res.Body)
}
