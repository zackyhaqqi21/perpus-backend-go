package route

import (
	"github.com/gin-gonic/gin"

	"perpusbackendgo/controller"
)

func RegisterBookRoutes(router *gin.Engine, ctrl *controller.BookController) {
	api := router.Group("/api/books")
	{
		api.GET("", ctrl.GetBooks)
		api.GET("/:id", ctrl.GetBookByID)
		api.POST("", ctrl.CreateBook)
		api.PUT("/:id", ctrl.UpdateBook)
		api.DELETE("/:id", ctrl.DeleteBook)
	}
}
