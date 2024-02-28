package model

// Album struct'ı
type Album struct {
	Base
	AlbumName   string   `json:"album_name"`
	Artist      string   `json:"artist"`
	ReleaseYear int      `json:"release_year"`
	Genre       string   `json:"genre"`
	TrackList   []string `json:"track_list"`
	CoverArt    string   `json:"cover_art"`
}
