package handlers

import (
	"go-cin/model"
	"go-cin/service"
	"go-cin/utils"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IAlbumHandler interface {
	CreateAlbum(ctx *gin.Context)
	GetAlbum(ctx *gin.Context)
	UpdateAlbum(ctx *gin.Context)
	DeleteAlbum(ctx *gin.Context)
	ListAlbums(ctx *gin.Context)
	UploadImage(ctx *gin.Context)
}

type AlbumHandler struct {
	service service.IAlbumService
	cld     *cloudinary.Cloudinary
}

func NewAlbumHandler(service service.IAlbumService, cld *cloudinary.Cloudinary) *AlbumHandler {
	return &AlbumHandler{
		service: service,
		cld:     cld,
	}
}

func (ah *AlbumHandler) UploadImage(ctx *gin.Context) {
	albumIDStr := ctx.Params.ByName("id")

	albumID, err := uuid.Parse(albumIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filename, ok := ctx.Get("filePath")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filename not found"})
	}

	file, ok := ctx.Get("file")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}
	imageUrl, err := utils.UploadToCloudinary(file.(multipart.File), filename.(string), ah.cld)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	album, err := ah.service.GetAlbum(albumID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	album.CoverArt = imageUrl
	err = ah.service.UpdateAlbum(albumID, album)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "Picture added succesfully!"})
}

func (ah *AlbumHandler) CreateAlbum(ctx *gin.Context) {
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

func (ah *AlbumHandler) DeleteAlbum(ctx *gin.Context) {
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

func (ah *AlbumHandler) UpdateAlbum(ctx *gin.Context) {
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

func (ah *AlbumHandler) GetAlbum(ctx *gin.Context) {
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

func (ah *AlbumHandler) ListAlbums(ctx *gin.Context) {
	albums, err := ah.service.ListAlbums()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": albums,
	})
}
