package router

import (
	"go-cin/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(handler *handlers.AlbumHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/", handler.CreateAlbum)
	r.GET("/:id", handler.GetAlbum)
	r.PUT("/:id", handler.UpdateAlbum)
	r.DELETE("/:id", handler.DeleteAlbum)
	r.GET("/", handler.ListAlbums)
	return r
}
