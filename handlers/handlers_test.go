package handlers

import (
	"go-cin/service"
	"reflect"
	"testing"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
)

func TestNewAlbumHandler(t *testing.T) {
	type args struct {
		service service.IAlbumService
		cld     *cloudinary.Cloudinary
	}
	tests := []struct {
		name string
		args args
		want *AlbumHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlbumHandler(tt.args.service, tt.args.cld); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlbumHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbumHandler_UploadImage(t *testing.T) {
	type fields struct {
		service service.IAlbumService
		cld     *cloudinary.Cloudinary
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ah := &AlbumHandler{
				service: tt.fields.service,
				cld:     tt.fields.cld,
			}
			ah.UploadImage(tt.args.ctx)
		})
	}
}

func TestAlbumHandler_CreateAlbum(t *testing.T) {
	type fields struct {
		service service.IAlbumService
		cld     *cloudinary.Cloudinary
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ah := &AlbumHandler{
				service: tt.fields.service,
				cld:     tt.fields.cld,
			}
			ah.CreateAlbum(tt.args.ctx)
		})
	}
}

func TestAlbumHandler_DeleteAlbum(t *testing.T) {
	type fields struct {
		service service.IAlbumService
		cld     *cloudinary.Cloudinary
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ah := &AlbumHandler{
				service: tt.fields.service,
				cld:     tt.fields.cld,
			}
			ah.DeleteAlbum(tt.args.ctx)
		})
	}
}

func TestAlbumHandler_UpdateAlbum(t *testing.T) {
	type fields struct {
		service service.IAlbumService
		cld     *cloudinary.Cloudinary
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ah := &AlbumHandler{
				service: tt.fields.service,
				cld:     tt.fields.cld,
			}
			ah.UpdateAlbum(tt.args.ctx)
		})
	}
}

func TestAlbumHandler_GetAlbum(t *testing.T) {
	type fields struct {
		service service.IAlbumService
		cld     *cloudinary.Cloudinary
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ah := &AlbumHandler{
				service: tt.fields.service,
				cld:     tt.fields.cld,
			}
			ah.GetAlbum(tt.args.ctx)
		})
	}
}

func TestAlbumHandler_ListAlbums(t *testing.T) {
	type fields struct {
		service service.IAlbumService
		cld     *cloudinary.Cloudinary
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ah := &AlbumHandler{
				service: tt.fields.service,
				cld:     tt.fields.cld,
			}
			ah.ListAlbums(tt.args.ctx)
		})
	}
}
