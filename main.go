package main

import (
	"database/sql"
	"log"
	"transactions_reader_stori/app_config/initializer"
	"transactions_reader_stori/app_config/initializer/controllers"
	"transactions_reader_stori/app_config/initializer/repositories"
	"transactions_reader_stori/app_config/initializer/routes"
	"transactions_reader_stori/app_config/initializer/services"
)

func main() {
	db := getDatabase()
	defer db.Close()

	appComponentsInitializer := initializer.AppComponentsInitializer{
		AppControllerFactoriesComponentsInitializer:  controllers.AppControllerFactoriesComponentsInitializer{},
		AppRepositoriesCommandsComponentsInitializer: repositories.AppRepositoriesCommandsComponentsInitializer{DB: db},
		AppServicesComponentsInitializer:             services.AppServicesComponentsInitializer{},
		AppRoutesInitializer:                         routes.RoutesInitializer{},
	}

	appComponentsInitializer.Init().Run()
}

func getDatabase() *sql.DB {
	db, err := sql.Open("sqlite", "transactions.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
