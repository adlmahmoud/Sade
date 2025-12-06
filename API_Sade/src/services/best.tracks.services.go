package services

import (
	"Tp_Api_spotify/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SearchTracksSrvices récupère les top tracks d'un artiste
// via l'endpoint /v1/artists/{id}/top-tracks.
func SearchTracksSrvices(artistID string) (*models.TopTracksResponse, int, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	url := "https://api.spotify.com/v1/artists/" + artistID + "/top-tracks?market=FR"

	request, requestErr := http.NewRequest(http.MethodGet, url, nil)
	if requestErr != nil {
		fmt.Printf("Erreur initialisiation requete - %s\n", requestErr.Error())
	}
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

	var topTracks models.TopTracksResponse
	decoderErr := json.NewDecoder(response.Body).Decode(&topTracks)
	if decoderErr != nil {
		fmt.Printf("Erreur decodage JSON - %s\n", decoderErr.Error())
	}
	return &topTracks, response.StatusCode, nil
}
