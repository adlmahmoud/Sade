package controllers

import (
	"Tp_Api_spotify/helpers"
	"Tp_Api_spotify/services"
	"Tp_Api_spotify/templates"
	"fmt"
	"net/http"
)

// DisplayTrack affiche les meilleurs titres (top tracks) d'un artiste.
// L'ID de l'artiste est passé en query param: /best-tracks?id=47zz7sob9NUcODy0BTDvKx
func DisplayTrack(w http.ResponseWriter, r *http.Request) {
	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		helpers.RedirectToError(w, r, http.StatusBadRequest, "Artist ID manquant")
		return
	}
	topTracks, status, err := services.SearchTracksSrvices(artistID)
	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des meilleurs titres")
		fmt.Println(err)
		return
	}
	// Le template s'appelle maintenant "best_track" et reçoit TopTracksResponse.
	templates.RenderTemplate(w, r, "best_track", topTracks)
}
