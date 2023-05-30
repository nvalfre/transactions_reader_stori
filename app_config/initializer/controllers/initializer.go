package controllers

import (
	"transactions_reader_stori/controllers/file_controller"
	"transactions_reader_stori/services/init_services"
)

type AppControllerFactoriesComponentsInitializerI interface {
	InitControllerFactories(services *init_services.Services) file_controller.FileControllerFactoryI
}

type AppControllerFactoriesComponentsInitializer struct {
}

func (initializer AppControllerFactoriesComponentsInitializer) InitControllerFactories(services *init_services.Services) file_controller.FileControllerFactoryI {
	fileControllerFactory := &file_controller.FileControllerFactory{FileService: services.FileService}
	return fileControllerFactory
}
