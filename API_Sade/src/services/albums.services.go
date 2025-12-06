package services

import (
	"Tp_Api_spotify/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SearchAlbumsSrvices(id_albums_sade string) (*models.AlbumsResponse, int, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	url := "https://api.spotify.com/v1/artists/" + id_albums_sade + "/albums"
	request, requestErr := http.NewRequest(http.MethodGet, url, nil)
	if requestErr != nil {
		fmt.Printf("Erreur initialisiation requete - %s\n", requestErr.Error())
	}
	// Récupération automatique du token Spotify
	token, tokenErr := GetSpotifyToken()
	if tokenErr != nil {
		fmt.Printf("Erreur récupération token Spotify - %s\n", tokenErr.Error())
	} else {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	response, responseErr := client.Do(request)
	if responseErr != nil {
		fmt.Printf("Erreur requete HTTP - %s\n", responseErr.Error())
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Erreur init requete - %d, status %s\n", response.StatusCode, response.Status)
	}

	var album models.AlbumsResponse
	decoderErr := json.NewDecoder(response.Body).Decode(&album)
	if decoderErr != nil {
		fmt.Printf("Erreur decodage JSON - %s\n", decoderErr.Error())
	}
	return &album, response.StatusCode, nil
}
