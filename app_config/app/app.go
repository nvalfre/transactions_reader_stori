package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

type appI interface {
	Run()
}

type App struct {
	router *gin.Engine
}

func (app App) Run() {
	if app.router == nil {
		log.Fatal("Invalid router engine")
	}
	if err := app.router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func NewApp(router *gin.Engine) App {
	return App{router}
}
