package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mariaafrv/gin-api/controller"
	"github.com/mariaafrv/gin-api/middlewares"
	"github.com/mariaafrv/gin-api/service"
)

var (
	filmService    service.FilmService       = service.New()
	filmController controller.FilmController = controller.New(filmService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/films", func(ctx *gin.Context) {
		ctx.JSON(200, filmController.FindAll())
	})

	server.POST("/films", func(ctx *gin.Context) {
		ctx.JSON(200, filmController.Save(ctx))
	})
	server.Run()
}
