package init_controllers

import (
	"transactions_reader_stori/controllers/file_controller"
	"transactions_reader_stori/services/init_services"
)

type Controllers struct {
	FileController *file_controller.FileController
}

func InitControllers(services *init_services.Services) *Controllers {
	fileController := file_controller.NewFileController(services.FileService)

	return &Controllers{
		FileController: fileController,
	}
}
