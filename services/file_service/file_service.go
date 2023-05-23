package file_service

import (
	"github.com/gin-gonic/gin"
	"log"
	"transactions_reader_stori/domain"
	"transactions_reader_stori/services/email_service"
)

// ProcessFile processes the uploaded file
func (h *FileService) ProcessFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to process file"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	fileContent := make([]byte, file.Size)
	_, err = src.Read(fileContent)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to read file content"})
		return
	}

	if h.execute(c, fileContent) {
		return
	}

	c.JSON(200, gin.H{"message": "File processed successfully"})
}

func (h *FileService) execute(c *gin.Context, fileContent []byte) bool {
	summaryCh := make(chan *domain.Summary)
	errorCh := make(chan error)
	doneCh := make(chan bool)

	go func(content []byte) {
		err := h.transactionService.ProcessFile(content)
		if err != nil {
			errorCh <- err
		} else {
			doneCh <- true
		}
	}(fileContent)

	go func() {
		summary, err := h.transactionService.GenerateSummary()
		if err != nil {
			errorCh <- err
		} else {
			summaryCh <- summary
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case summary := <-summaryCh:
			emailService := email_service.NewEmailServiceDefault()
			err := emailService.SendSummaryEmail(summary, "nvalfre@gmail.com")
			if err != nil {
				log.Println("Failed to send summary email:", err)
				errorCh <- err
				return true
			}
		case err := <-errorCh:
			log.Println("Error occurred:", err)
			c.JSON(500, gin.H{"error": "Failed to process file"})
			return true
		case <-doneCh:
			log.Println("File processing completed successfully")
		}
	}
	return false
}
