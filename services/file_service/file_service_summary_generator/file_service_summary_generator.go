package file_service_summary_generator

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"transactions_reader_stori/domain"
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/transaction_service"
)

const AccountId = "account_id"
const AccountName = "account_name"

type FileSummaryGeneratorUseCase struct {
	transactionService transaction_service.TransactionServiceI
	emailService       email_service.EmailServiceI
}

func (fs *FileSummaryGeneratorUseCase) GenerateSummary(c *gin.Context, fileContent []byte) (*domain.SummaryVO, error) {
	summaryCh := make(chan *domain.SummaryVO)
	errorCh := make(chan error)
	fileProcessingDoneCh := make(chan bool)
	summaryGenerationDoneCh := make(chan bool)

	accountID := parseIntParam(c.Query("account_id"))
	accountName := c.Query("name")
	accountEmail := c.Query("email")

	go func(content []byte) {
		err := fs.transactionService.ProcessFileContent(content, accountID, accountName, accountEmail)
		if err != nil {
			errorCh <- err
		} else {
			fileProcessingDoneCh <- true
		}
	}(fileContent)

	go func() {
		select {
		case <-fileProcessingDoneCh:
			log.Println("File processing completed successfully")
			summary, err := fs.transactionService.GenerateSummary(accountID)
			if err != nil {
				errorCh <- err
			} else {
				summaryCh <- summary
			}
		case err := <-errorCh:
			log.Println("Error occurred on file processing:", err)
			errorCh <- err
		}
		close(summaryGenerationDoneCh)
	}()

	var summary *domain.SummaryVO
	select {
	case summary = <-summaryCh:
		err := fs.emailService.SendSummaryEmail(summary, "testmailnv23@gmail.com")
		if err != nil {
			log.Println("Failed to send summary email:", err)
			errorCh <- err
			return nil, err
		}
	case err := <-errorCh:
		log.Println("Error occurred on summary send:", err)
		return nil, err
	}

	<-summaryGenerationDoneCh

	return summary, nil
}

func parseIntParam(string string) int {
	i, err := strconv.ParseInt(string, 10, 0)
	if err != nil {
		log.Fatal("Failed to read account id:", err)
	}
	return int(i)
}
