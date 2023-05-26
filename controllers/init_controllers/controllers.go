package init_controllers

import (
	"transactions_reader_stori/controllers/file_controller"
)

type Controllers struct {
	FileController *file_controller.FileController
}

func InitWith(fileControllerFactory file_controller.FileControllerFactoryI) *Controllers {
	return &Controllers{
		FileController: fileControllerFactory.NewFileController(),
	}
}
