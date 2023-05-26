package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"transactions_reader_stori/controllers"
	"transactions_reader_stori/repository"
	"transactions_reader_stori/repository/account_repository"
	"transactions_reader_stori/repository/transaction_repository"
	"transactions_reader_stori/services"
)

func RunApp() {
	databaseRepo := initDb()

	transactionDatabaseRepo := transaction_repository.NewTransactionDatabaseRepo(databaseRepo)
	accountDatabaseRepo := account_repository.NewAccountDatabaseRepo(databaseRepo)

	initServices := services.InitServices(transactionDatabaseRepo, accountDatabaseRepo)
	initControllers := controllers.InitControllers(initServices)
	router := initRoutes(initControllers)

	run(router)
}

func run(router *gin.Engine) {
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initRoutes(controllers *controllers.Controllers) *gin.Engine {
	router := gin.Default()

	router.POST("/file/process/transactions", controllers.FileController.ProcessFile)

	return router
}

func initDb() *repository.DatabaseRepo {
	// db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }
	// db.AutoMigrate(&dao.Account{}, &dao.Transaction{})
	// repository.NewDatabaseRepo(db)

	return repository.NewDatabaseRepo()
}
