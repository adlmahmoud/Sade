package models

type AlbumsResponse struct {
	Items  []AlbumModel `json:"items"`
	Total  int          `json:"total"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
}

type AlbumModel struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	AlbumType   string  `json:"album_type"`
	TotalTracks int     `json:"total_tracks"`
	ReleaseDate string  `json:"release_date"`
	ReleasePrec string  `json:"release_date_precision"`
	Images      []Image `json:"images"`
	Href        string  `json:"href"`
}
