package router

import (
	"go-cin/handlers"
	"go-cin/middleware"

	"github.com/gin-gonic/gin"
)

type router struct {
	albumHandler handlers.IAlbumHandler
}

func NewRouter(albumHandler handlers.IAlbumHandler) *router {
	return &router{
		albumHandler: albumHandler,
	}

}

func (h *router) SetUpRouter(r *gin.Engine) *gin.Engine {

	apiAlbum := r.Group("/api/album")
	{
		apiAlbum.POST("/", h.albumHandler.CreateAlbum)
		apiAlbum.GET("/:id", h.albumHandler.GetAlbum)
		apiAlbum.PUT("/:id", h.albumHandler.UpdateAlbum)
		apiAlbum.DELETE("/:id", h.albumHandler.DeleteAlbum)
		apiAlbum.GET("/", h.albumHandler.ListAlbums)
		apiAlbum.PUT("/update-pic/:id", middleware.FileUploadMiddleware(), h.albumHandler.UploadImage)
	}

	return r
}
