package handlers

import (
	"go-cin/model"
	"go-cin/service"
	"net/http"

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

type albumHandler struct {
	service service.IAlbumService
}

func NewAlbumHandler(service service.IAlbumService) IAlbumHandler {
	return &albumHandler{service: service}
}

func (ah *albumHandler) CreateAlbum(ctx *gin.Context) {
	var album model.Album

	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ah.service.CreateAlbum(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "album created successfully"})
}

func (ah *albumHandler) DeleteAlbum(ctx *gin.Context) {
	albumIDStr := ctx.Params.ByName("id")

	albumID, err := uuid.Parse(albumIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ah.service.DeleteAlbum(albumID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "album deleted successfully"})
}

func (ah *albumHandler) UpdateAlbum(ctx *gin.Context) {
	albumIDStr := ctx.Params.ByName("id")
	albumID, err := uuid.Parse(albumIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var album model.Album
	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ah.service.UpdateAlbum(albumID, &album); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "album updated successfully"})
}

func (ah *albumHandler) GetAlbum(ctx *gin.Context) {
	albumIDStr := ctx.Params.ByName("id")
	albumID, err := uuid.Parse(albumIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	album, err := ah.service.GetAlbum(albumID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": album,
	})
}

func (ah *albumHandler) ListAlbums(ctx *gin.Context) {
	albums, err := ah.service.ListAlbums()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": albums,
	})
}
