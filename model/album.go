package model

// Album struct'ı
type Album struct {
	Base
	AlbumName   string `json:"album_name"`
	Artist      string `json:"artist"`
	ReleaseYear int    `json:"release_year"`
	Genre       string `json:"genre"`
	CoverArt    string `json:"cover_art"`
}
