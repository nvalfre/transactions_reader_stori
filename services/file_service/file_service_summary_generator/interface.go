package file_service_summary_generator

import (
	"github.com/gin-gonic/gin"
	"transactions_reader_stori/domain"
)

type FileSummaryGeneratorUseCaseI interface {
	Execute(c *gin.Context, fileContent []byte) (*domain.SummaryVO, error)
}
