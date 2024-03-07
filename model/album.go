package model

// Album struct'ı
type Album struct {
	Base
	AlbumName   string `json:"album_name" example:"Abbey Road"`
	Artist      string `json:"artist" example:"The Beatles"`
	ReleaseYear int    `json:"release_year" example:"1969"`
	Genre       string `json:"genre" example:"Rock"`
	CoverArt    string `json:"cover_art" example:"https://example.com/abbey_road_cover.jpg"`
}
