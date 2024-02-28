package repository

import (
	"go-cin/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAlbum interface {
	CreateAlbum(album *model.Album) error
	UpdateAlbum(album *model.Album) error
	DeleteAlbum(id uuid.UUID) error
	GetAlbum(id uuid.UUID) (*model.Album, error)
	GetAlbumByName(name string) (*model.Album, error)
	GetAlbumByArtist(artist string) (*model.Album, error)
	ListAlbums() ([]*model.Album, error)
}

type AlbumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (repo *AlbumRepository) CreateAlbum(album *model.Album) error {
	return repo.db.Create(album).Error
}

func (repo *AlbumRepository) UpdateAlbum(album *model.Album) error {
	return repo.db.Save(album).Error
}

func (repo *AlbumRepository) DeleteAlbum(id uuid.UUID) error {
	return repo.db.Where("id=?").Delete(id).Error
}

func (repo *AlbumRepository) GetAlbum(id uuid.UUID) (*model.Album, error) {
	var album model.Album
	err := repo.db.Where("id =?", id).First(&album).Error
	return &album, err
}

func (repo *AlbumRepository) GetAlbumByName(name string) (*model.Album, error) {
	var album model.Album
	err := repo.db.Where("album_name =?", name).First(&album).Error
	return &album, err
}

func (repo *AlbumRepository) GetAlbumByArtist(artist string) (*model.Album, error) {
	var album model.Album
	err := repo.db.Where("artist =?", artist).First(&album).Error
	return &album, err
}

func (repo *AlbumRepository) ListAlbums() ([]*model.Album, error) {
	var albums []*model.Album
	err := repo.db.Find(&albums).Error
	return albums, err
}
