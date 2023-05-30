package routes

import (
	"github.com/gin-gonic/gin"
	"transactions_reader_stori/controllers/init_controllers"
)

type RoutesInitializerI interface {
	InitRoutes(controllers *init_controllers.Controllers) *gin.Engine
}

type RoutesInitializer struct {
}

func (initializer RoutesInitializer) InitRoutes(controllers *init_controllers.Controllers) *gin.Engine {
	router := gin.Default()

	router.POST("/file/process/transactions", controllers.FileController.ProcessFile)

	return router
}
