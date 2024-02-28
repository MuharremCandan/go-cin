package service

import (
	"go-cin/model"
	"go-cin/repository"
)

type IAlbumService interface {
	CreateAlbum(album *model.Album) error
	UpdateAlbum(album *model.Album) error
	DeleteAlbum(album *model.Album) error
	ListAlbums() ([]*model.Album, error) //
}

type albumService struct {
	repo repository.IAlbum
}

func NewAlbumService(repo repository.IAlbum) *albumService {
	return &albumService{repo: repo}
}

func (s *albumService) CreateAlbum(album *model.Album) error {
	return s.repo.CreateAlbum(album)
}

func (s *albumService) UpdateAlbum(album *model.Album) error {
	return s.repo.UpdateAlbum(album)
}

func (s *albumService) DeleteAlbum(album *model.Album) error {
	return s.repo.DeleteAlbum(album)
}

func (s *albumService) ListAlbums() ([]*model.Album, error) {
	return s.repo.ListAlbums()
}
