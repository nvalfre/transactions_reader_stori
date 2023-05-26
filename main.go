package main

import (
	"transactions_reader_stori/app_config/initializer"
)

func main() {
	appComponentsInitializer := initializer.AppComponentsInitializer{}
	appComponentsInitializer.Init()
}
