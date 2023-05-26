package main

import (
	"transactions_reader_stori/main/initializer"
)

func main() {
	appComponentsInitializer := initializer.AppComponentsInitializer{}
	appComponentsInitializer.Init()
}
