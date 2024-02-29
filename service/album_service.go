package service

import (
	"fmt"
	"go-cin/model"
	"go-cin/repository"

	"github.com/google/uuid"
)

type IAlbumService interface {
	CreateAlbum(album *model.Album) error
	UpdateAlbum(id uuid.UUID, album *model.Album) error
	DeleteAlbum(id uuid.UUID) error
	ListAlbums() ([]*model.Album, error)
	GetAlbum(id uuid.UUID) (*model.Album, error)
}

type albumService struct {
	repo repository.IAlbum
}

func NewAlbumService(repo repository.IAlbum) IAlbumService {
	return &albumService{
		repo: repo,
	}
}

func (s *albumService) CreateAlbum(album *model.Album) error {
	return s.repo.CreateAlbum(album)
}

func (s *albumService) UpdateAlbum(id uuid.UUID, album *model.Album) error {
	return s.repo.UpdateAlbum(id, album)
}

func (s *albumService) DeleteAlbum(id uuid.UUID) error {
	_, err := s.repo.GetAlbum(id)
	if err != nil {
		return fmt.Errorf("record not found: %v", err)
	}
	return s.repo.DeleteAlbum(id)
}

func (s *albumService) ListAlbums() ([]*model.Album, error) {
	return s.repo.ListAlbums()
}

func (s *albumService) GetAlbum(id uuid.UUID) (*model.Album, error) {
	return s.repo.GetAlbum(id)
}
