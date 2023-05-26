package initializer

import (
	"github.com/gin-gonic/gin"
	"transactions_reader_stori/controllers/file_controller"
	"transactions_reader_stori/controllers/init_controllers"
	"transactions_reader_stori/main/app"
	"transactions_reader_stori/repository/factories/account_repository_factory"
	"transactions_reader_stori/repository/factories/transaction_repository_factory"
	"transactions_reader_stori/repository/init_repositories"
	"transactions_reader_stori/services/account_service"
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/file_service"
	"transactions_reader_stori/services/init_services"
	"transactions_reader_stori/services/transaction_service"
)

type appComponentsInitializerI interface {
	Init()
	initDatabaseRepoCommands() *init_repositories.DatabaseRepoCommands
	initControllerFactories(services *init_services.Services) file_controller.FileControllerFactoryI
	initServicesFactories() (*account_service.AccountServiceFactory, *transaction_service.TransactionServiceFactory, *file_service.FileServiceFactory, *email_service.EmailServiceFactory)
	initRoutes(controllers *init_controllers.Controllers) *gin.Engine
}

type AppComponentsInitializer struct {
}

func (initializer AppComponentsInitializer) Init() {
	databaseRepoCommands := initializer.initDatabaseRepoCommands()

	accountServiceFactory, transactionServiceFactory, fileServiceFactory, emailServiceFactory := initializer.initServicesFactories()

	services := init_services.InitWith(
		accountServiceFactory,
		transactionServiceFactory,
		fileServiceFactory,
		emailServiceFactory,
		databaseRepoCommands.TransactionRepository,
		databaseRepoCommands.AccountRepository,
	)

	fileControllerFactory := initializer.initControllerFactories(services)

	controllers := init_controllers.InitWith(fileControllerFactory)
	routes := initializer.initRoutes(controllers)
	app.NewApp(routes).Run()
}

func (initializer AppComponentsInitializer) initDatabaseRepoCommands() *init_repositories.DatabaseRepoCommands {
	repo := init_repositories.NewDatabaseRepo()

	transactionDatabaseRepo := transaction_repository_factory.NewTransactionDatabaseRepo(repo)
	accountDatabaseRepo := account_repository_factory.NewAccountDatabaseRepo(repo)

	commands := &init_repositories.DatabaseRepoCommands{
		TransactionRepository: transactionDatabaseRepo,
		AccountRepository:     accountDatabaseRepo,
	}

	return commands
}

func (initializer AppComponentsInitializer) initControllerFactories(services *init_services.Services) file_controller.FileControllerFactoryI {
	fileControllerFactory := &file_controller.FileControllerFactory{FileService: services.FileService}
	return fileControllerFactory
}

func (initializer AppComponentsInitializer) initServicesFactories() (*account_service.AccountServiceFactory, *transaction_service.TransactionServiceFactory, *file_service.FileServiceFactory, *email_service.EmailServiceFactory) {
	accountServiceFactory := &account_service.AccountServiceFactory{}
	transactionServiceFactory := &transaction_service.TransactionServiceFactory{}
	fileServiceFactory := &file_service.FileServiceFactory{}
	emailServiceFactory := &email_service.EmailServiceFactory{}
	return accountServiceFactory, transactionServiceFactory, fileServiceFactory, emailServiceFactory
}

func (initializer AppComponentsInitializer) initRoutes(controllers *init_controllers.Controllers) *gin.Engine {
	router := gin.Default()

	router.POST("/file/process/transactions", controllers.FileController.ProcessFile)

	return router
}
