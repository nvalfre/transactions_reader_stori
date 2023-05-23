package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"transactions_reader_stori/domain/dao"
	"transactions_reader_stori/repository"
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/file_service"
	"transactions_reader_stori/services/transaction_service"
)

type Services struct {
	transactionService *transaction_service.TransactionService
	fileService        *file_service.FileService
	emailService       *email_service.EmailService
}

func RunApp() {
	databaseRepo := initDb()

	services := initServices(databaseRepo)

	router := initRoutes(services)

	run(router)
}

func run(router *gin.Engine) {
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initRoutes(services Services) *gin.Engine {

	router := gin.Default()
	router.POST("/file/process/transactions", services.fileService.ProcessFile)
	return router
}

func initServices(db *repository.DatabaseRepo) Services {
	transactionService := transaction_service.NewTransactionService(db)
	emailService := email_service.NewEmailServiceDefault()
	fileService := file_service.NewFileService(transactionService, emailService)

	return Services{
		transactionService: transactionService,
		fileService:        fileService,
		emailService:       emailService,
	}
}

func initDb() *repository.DatabaseRepo {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&dao.Account{}, &dao.Transaction{})

	databaseRepo := repository.NewDatabaseRepo(db)
	return databaseRepo
}
