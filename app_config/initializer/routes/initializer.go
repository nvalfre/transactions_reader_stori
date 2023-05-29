package routes

import (
	"github.com/gin-gonic/gin"
	"transactions_reader_stori/controllers/init_controllers"
)

type RoutesInitializerI interface {
	InitRoutes(controllers *init_controllers.Controllers) *gin.Engine
}

type RoutesInitializer struct {
	engine *gin.Engine
}

func (routesInitializer RoutesInitializer) InitRoutes(controllers *init_controllers.Controllers) *gin.Engine {
	router := gin.Default()

	routesInitializer.init(controllers, router)

	return router
}
