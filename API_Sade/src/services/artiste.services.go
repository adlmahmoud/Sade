package services

import (
	"Tp_Api_spotify/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func SearchArtisteSrvices(id string) (*models.ArtistModel, int, error) {
	clients := http.Client{
		Timeout: 5 * time.Second,
	}
	url := fmt.Sprintf("https://api.spotify.com/v1/artists/%v", id)
	request, requestErr := http.NewRequest(http.MethodGet, url, nil)
	if requestErr != nil {
		fmt.Printf("Erreur init requete - %s\n", requestErr.Error())
	}
	token := os.Getenv("SPOTIFY_TOKEN")
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	response, responseErr := clients.Do(request)
	if responseErr != nil {
		fmt.Printf("Erreur requete HTTP - %s\n", responseErr.Error())
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Erreur init requete - %d, status %s\n", response.StatusCode, response.Status)
	}

	var artiste models.ArtistModel
	decoderErr := json.NewDecoder(response.Body).Decode(&artiste)
	if decoderErr != nil {
		fmt.Printf("Erreur decodage JSON - %s\n", decoderErr.Error())
	}
	return &artiste, response.StatusCode, nil
}
