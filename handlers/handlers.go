package handlers

import (
	"go-cin/model"
	"go-cin/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IAlbumHandler interface {
	CreateAlbum(ctx *gin.Context)
	GetAlbum(ctx *gin.Context)
	UpdateAlbum(ctx *gin.Context)
	DeleteAlbum(ctx *gin.Context)
	ListAlbums(ctx *gin.Context)
}

type AlbumHandler struct {
	service *service.AlbumService
}

func NewAlbumHandler(service *service.AlbumService) *AlbumHandler {
	return &AlbumHandler{service: service}
}

func (ah *AlbumHandler) CreateAlbum(ctx *gin.Context) {
	var album model.Album

	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	if err := ah.service.CreateAlbum(&album); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"success": "album created successfully"})
}

func (ah *AlbumHandler) DeleteAlbum(ctx *gin.Context) {
	albumIDStr := ctx.Params.ByName("id")

	albumID, err := uuid.Parse(albumIDStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	if err := ah.service.DeleteAlbum(albumID); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, gin.H{"success": "album deleted successfully"})
}

func (ah *AlbumHandler) UpdateAlbum(ctx *gin.Context) {
	var album model.Album
	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	if err := ah.service.UpdateAlbum(&album); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, gin.H{"success": "album updated successfully"})
}

func (ah *AlbumHandler) GetAlbum(ctx *gin.Context) {
	albumIDStr := ctx.Params.ByName("id")
	albumID, err := uuid.Parse(albumIDStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	album, err := ah.service.GetAlbum(albumID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{
		"success": album,
	})
}

func (ah *AlbumHandler) ListAlbums(ctx *gin.Context) {
	albums, err := ah.service.ListAlbums()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{
		"success": albums,
	})
}
