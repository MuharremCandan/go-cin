package handlers

import (
	"go-cin/model"
	"go-cin/pkg/utils"
	"go-cin/service"
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

// @Summary Upload an image
// @Description Upload an image to the server
// @Tags images
// @ID upload-image
// @Accept mpfd
// @Produce json
// @Param file formData file true "Image file to upload"
// @Success 200
// @Router /upload/image/:id [post]
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

// @Summary Create a new album
// @Description Create a new album with the provided details
// @Tags albums
// @ID create-album
// @Accept json
// @Produce json
// @Param album body model.Album true "Album details"
// @Success 200 {string} string "Successfully created album"
// @Failure 400 {object} string "Bad Request"
// @Router /album [post]
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

// @Summary Delete an album by ID
// @Description Delete an album by its ID
// @Tags albums
// @ID delete-album-by-id
// @Produce json
// @Param id path int true "Album ID"
// @Success 204 "No Content"
// @Router /album/:id [delete]
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

// @Summary Update an album by ID
// @Description Update an existing album with the provided details
// @Tags albums
// @ID update-album-by-id
// @Produce json
// @Param id path int true "Album ID"
// @Param album body model.Album true "Updated album details"
// @Success 200 {object} model.Album
// @Router /albums/:id [put]
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

// @Summary Get an album by ID
// @Description Get an album details by its ID
// @Tags albums
// @ID get-album-by-id
// @Produce json
// @Param id path int true "Album ID"
// @Success 200 {object} model.Album
// @Router /album/:id [get]
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

// @Summary Get all albums
// @Description Get a list of all albums
// @Tags albums
// @ID get-all-albums
// @Produce json
// @Success 200 {array} model.Album
// @Router /album [get]
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
