package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const spotifyTokenURL = "https://accounts.spotify.com/api/token"

// spotifyTokenResponse représente la réponse de l'API /api/token de Spotify.
type spotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

var (
	spotifyToken     string
	spotifyTokenExp  time.Time
	spotifyTokenLock sync.RWMutex
)

// GetSpotifyToken retourne un token d'accès Spotify valide.
// Si un token est déjà en cache et encore valable, il est réutilisé.
// Sinon, on en demande un nouveau via le client credentials flow.
func GetSpotifyToken() (string, error) {
	spotifyTokenLock.RLock()
	if spotifyToken != "" && time.Now().Before(spotifyTokenExp.Add(-30*time.Second)) {
		t := spotifyToken
		spotifyTokenLock.RUnlock()
		return t, nil
	}
	spotifyTokenLock.RUnlock()

	return refreshSpotifyToken()
}

// refreshSpotifyToken demande un nouveau token d'accès à Spotify et le met en cache.
func refreshSpotifyToken() (string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		return "", fmt.Errorf("SPOTIFY_CLIENT_ID ou SPOTIFY_CLIENT_SECRET non configurés")
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, spotifyTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("erreur création requête token Spotify: %w", err)
	}

	// Authentification Basic base64(client_id:client_secret)
	credentials := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	req.Header.Set("Authorization", "Basic "+credentials)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erreur requête token Spotify: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erreur token Spotify: status %d (%s)", resp.StatusCode, resp.Status)
	}

	var tokResp spotifyTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokResp); err != nil {
		return "", fmt.Errorf("erreur décodage réponse token Spotify: %w", err)
	}

	if tokResp.AccessToken == "" {
		return "", fmt.Errorf("réponse token Spotify sans access_token")
	}

	spotifyTokenLock.Lock()
	defer spotifyTokenLock.Unlock()

	spotifyToken = tokResp.AccessToken
	spotifyTokenExp = time.Now().Add(time.Duration(tokResp.ExpiresIn) * time.Second)

	return spotifyToken, nil
}
