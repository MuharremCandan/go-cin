package router

import (
	"go-cin/handlers"
	"go-cin/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(handler handlers.IAlbumHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/", handler.CreateAlbum)
	r.GET("/:id", handler.GetAlbum)
	r.PUT("/:id", handler.UpdateAlbum)
	r.DELETE("/:id", handler.DeleteAlbum)
	r.GET("/", handler.ListAlbums)
	r.PUT("/update-pic/:id", middleware.FileUploadMiddleware(), handler.UploadImage)
	return r
}
