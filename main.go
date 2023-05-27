package main

import (
	"transactions_reader_stori/app_config/initializer"
	"transactions_reader_stori/app_config/initializer/controllers"
	"transactions_reader_stori/app_config/initializer/repositories"
	"transactions_reader_stori/app_config/initializer/routes"
	"transactions_reader_stori/app_config/initializer/services"
)

func main() {
	appComponentsInitializer := initializer.AppComponentsInitializer{
		AppControllerFactoriesComponentsInitializer:  controllers.AppControllerFactoriesComponentsInitializer{},
		AppRepositoriesCommandsComponentsInitializer: repositories.AppRepositoriesCommandsComponentsInitializer{},
		AppServicesComponentsInitializer:             services.AppServicesComponentsInitializer{},
		AppRoutesInitializer:                         routes.RoutesInitializer{},
	}
	appComponentsInitializer.Init()
}
