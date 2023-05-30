package initializer

import (
	"transactions_reader_stori/app_config/app"
	"transactions_reader_stori/app_config/initializer/controllers"
	"transactions_reader_stori/app_config/initializer/repositories"
	"transactions_reader_stori/app_config/initializer/routes"
	"transactions_reader_stori/app_config/initializer/services"
	"transactions_reader_stori/controllers/init_controllers"
	"transactions_reader_stori/services/init_services"
)

type appComponentsInitializerI interface {
	Init() app.App
}

type AppComponentsInitializer struct {
	AppControllerFactoriesComponentsInitializer  controllers.AppControllerFactoriesComponentsInitializerI
	AppRepositoriesCommandsComponentsInitializer repositories.AppRepositoriesCommandsComponentsInitializerI
	AppServicesComponentsInitializer             services.AppServicesComponentsInitializerI
	AppRoutesInitializer                         routes.RoutesInitializerI
}

func (initializer AppComponentsInitializer) Init() app.App {
	databaseRepoCommands := initializer.AppRepositoriesCommandsComponentsInitializer.InitDatabaseRepoCommands()

	accountServiceFactory, transactionServiceFactory, fileServiceFactory, emailServiceFactory := initializer.AppServicesComponentsInitializer.InitServicesFactories()

	appServices := init_services.InitWith(
		accountServiceFactory,
		transactionServiceFactory,
		fileServiceFactory,
		emailServiceFactory,
		databaseRepoCommands.TransactionRepository,
		databaseRepoCommands.AccountRepository,
	)

	fileControllerFactory := initializer.AppControllerFactoriesComponentsInitializer.InitControllerFactories(appServices)

	appControllers := init_controllers.InitWith(fileControllerFactory)
	appRoutes := initializer.AppRoutesInitializer.InitRoutes(appControllers)
	return app.NewApp(appRoutes)
}
