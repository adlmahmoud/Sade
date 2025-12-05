package models

// TopTracksResponse représente la réponse de /v1/artists/{id}/top-tracks
type TopTracksResponse struct {
	Tracks []Track `json:"tracks"`
}

type Track struct {
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	DurationMs       int            `json:"duration_ms"`
	Artists          []SimpleArtist `json:"artists"`
	Album            SimpleAlbum    `json:"album"`
	Popularity       int            `json:"popularity"`
	PreviewURL       string         `json:"preview_url"`
	TrackNumber      int            `json:"track_number"`
	DiscNumber       int            `json:"disc_number"`
	AvailableMarkets []string       `json:"available_markets"`
	Href             string         `json:"href"`
	URI              string         `json:"uri"`
}

type SimpleAlbum struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}

type SimpleArtist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
	URI  string `json:"uri"`
}
