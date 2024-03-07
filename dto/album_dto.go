// dto.go

package dto

// CreateAlbumRequest represents the request structure for creating an album.
type CreateAlbumRequest struct {
	AlbumName   string `json:"album_name" binding:"required" example:"Abbey Road"`
	Artist      string `json:"artist" binding:"required" example:"The Beatles"`
	ReleaseYear int    `json:"release_year" binding:"required" example:"1969"`
	Genre       string `json:"genre" binding:"required" example:"Rock"`
	CoverArt    string `json:"cover_art" binding:"required" example:"https://example.com/abbey_road_cover.jpg"`
}

// CreateAlbumResponse represents the response structure for a successfully created album.
type CreateAlbumResponse struct {
	H map[string]interface{} `json:"h" binding:"required"`
}
