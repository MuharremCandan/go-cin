package router

import (
	"go-cin/handlers"
	"go-cin/middleware"

	docs "go-cin/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	// Swagger UI için
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		apiAlbum := v1.Group("/album")
		{
			apiAlbum.POST("/", h.albumHandler.CreateAlbum)
			apiAlbum.GET("/:id", h.albumHandler.GetAlbum)
			apiAlbum.PUT("/:id", h.albumHandler.UpdateAlbum)
			apiAlbum.DELETE("/:id", h.albumHandler.DeleteAlbum)
			apiAlbum.GET("/", h.albumHandler.ListAlbums)
			apiAlbum.POST("/upload/iamge/:id", middleware.FileUploadMiddleware(), h.albumHandler.UploadImage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
