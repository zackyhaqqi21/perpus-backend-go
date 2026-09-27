package main

import (
	"github.com/gin-gonic/gin"

	"perpusbackendgo/config"
	"perpusbackendgo/controller"
	"perpusbackendgo/repository"
	"perpusbackendgo/route"
	"perpusbackendgo/service"
)

func main() {
	config.ConnectDatabase()

	bookRepo := repository.NewBookRepository(config.DB)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	router := gin.Default()
	route.RegisterBookRoutes(router, bookController)

	router.Run(":8080")
}
