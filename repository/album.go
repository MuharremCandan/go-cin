package repository

import (
	"go-cin/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAlbum interface {
	CreateAlbum(album *model.Album) error
	UpdateAlbum(id uuid.UUID, album *model.Album) error
	DeleteAlbum(id uuid.UUID) error
	GetAlbum(id uuid.UUID) (*model.Album, error)
	GetAlbumByName(name string) (*model.Album, error)
	GetAlbumByArtist(artist string) (*model.Album, error)
	ListAlbums() ([]*model.Album, error)
}

type albumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) IAlbum {
	return &albumRepository{
		db: db,
	}
}

func (repo *albumRepository) CreateAlbum(album *model.Album) error {
	return repo.db.Create(album).Error
}

func (repo *albumRepository) UpdateAlbum(id uuid.UUID, album *model.Album) error {
	return repo.db.Where("id=?", id).Save(album).Error
}

func (repo *albumRepository) DeleteAlbum(id uuid.UUID) error {
	return repo.db.Where("id=?", id).Delete(&model.Album{}).Error
}

func (repo *albumRepository) GetAlbum(id uuid.UUID) (*model.Album, error) {
	var album model.Album
	err := repo.db.Where("id =?", id).First(&album).Error
	return &album, err
}

func (repo *albumRepository) GetAlbumByName(name string) (*model.Album, error) {
	var album model.Album
	err := repo.db.Where("album_name =?", name).First(&album).Error
	return &album, err
}

func (repo *albumRepository) GetAlbumByArtist(artist string) (*model.Album, error) {
	var album model.Album
	err := repo.db.Where("artist =?", artist).First(&album).Error
	return &album, err
}

func (repo *albumRepository) ListAlbums() ([]*model.Album, error) {
	var albums []*model.Album
	err := repo.db.Find(&albums).Error
	return albums, err
}
