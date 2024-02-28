package service

import (
	"go-cin/model"
	"go-cin/repository"

	"github.com/google/uuid"
)

type IAlbumService interface {
	CreateAlbum(album *model.Album) error
	UpdateAlbum(album *model.Album) error
	DeleteAlbum(id uuid.UUID) error
	ListAlbums() ([]*model.Album, error)
	GetAlbum(id uuid.UUID) (*model.Album, error)
}

type AlbumService struct {
	repo *repository.AlbumRepository
}

func NewAlbumService(repo *repository.AlbumRepository) *AlbumService {
	return &AlbumService{repo: repo}
}

func (s *AlbumService) CreateAlbum(album *model.Album) error {
	return s.repo.CreateAlbum(album)
}

func (s *AlbumService) UpdateAlbum(album *model.Album) error {
	return s.repo.UpdateAlbum(album)
}

func (s *AlbumService) DeleteAlbum(id uuid.UUID) error {
	return s.repo.DeleteAlbum(id)
}

func (s *AlbumService) ListAlbums() ([]*model.Album, error) {
	return s.repo.ListAlbums()
}

func (s *AlbumService) GetAlbum(id uuid.UUID) (*model.Album, error) {
	return s.repo.GetAlbum(id)
}
