package models

type ArtistModel struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Genres     []string `json:"genres"`
	Popularity int      `json:"popularity"`
	Images     []Image  `json:"images"`
	Type       string   `json:"type"`
	Href       string   `json:"href"`
}
