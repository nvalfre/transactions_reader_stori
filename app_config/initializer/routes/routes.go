package routes

import (
	"github.com/gin-gonic/gin"
	"transactions_reader_stori/controllers/init_controllers"
)

func (routesInitializer RoutesInitializer) init(controllers *init_controllers.Controllers, router *gin.Engine) {
	router.POST(fileProcessTransactionsPath, controllers.FileController.ProcessFile)
}
