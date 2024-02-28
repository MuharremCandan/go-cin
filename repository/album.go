package repository

import (
	"go-cin/model"

	"gorm.io/gorm"
)

type IAlbum interface {
	CreateAlbum(album *model.Album) error
	UpdateAlbum(album *model.Album) error
	DeleteAlbum(album *model.Album) error
	GetAlbum(id int) (*model.Album, error)
	GetAlbumByName(name string) (*model.Album, error)
	GetAlbumByArtist(artist string) (*model.Album, error)
	ListAlbums() ([]*model.Album, error)
}

type albumrepository struct {
	db *gorm.DB
}

func NewAlbumrepository(db *gorm.DB) *albumrepository {
	return &albumrepository{db: db}
}

func (repo *albumrepository) CreateAlbum(album *model.Album) error {
	return repo.db.Create(album).Error
}

func (repo *albumrepository) UpdateAlbum(album *model.Album) error {
	return repo.db.Save(album).Error
}

func (repo *albumrepository) DeleteAlbum(album *model.Album) error {
	return repo.db.Delete(album).Error
}

func (repo *albumrepository) GetAlbum(id int) (*model.Album, error) {
	var album model.Album
	err := repo.db.Where("id =?", id).First(&album).Error
	return &album, err
}

func (repo *albumrepository) GetAlbumByName(name string) (*model.Album, error) {
	var album model.Album
	err := repo.db.Where("album_name =?", name).First(&album).Error
	return &album, err
}

func (repo *albumrepository) GetAlbumByArtist(artist string) (*model.Album, error) {
	var album model.Album
	err := repo.db.Where("artist =?", artist).First(&album).Error
	return &album, err
}

func (repo *albumrepository) ListAlbums() ([]*model.Album, error) {
	var albums []*model.Album
	err := repo.db.Find(&albums).Error
	return albums, err
}
