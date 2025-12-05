package controllers

import (
	"Tp_Api_spotify/helpers"
	"Tp_Api_spotify/services"
	"Tp_Api_spotify/templates"
	"fmt"
	"net/http"
)

func DisplayListAlbums(w http.ResponseWriter, r *http.Request) {
	id_albums_sade := "47zz7sob9NUcODy0BTDvKx"
	albums, status, err := services.SearchAlbumsSrvices(id_albums_sade)
	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des albums")
		fmt.Println(err)
		return
	}
	templates.RenderTemplate(w, r, "albums", albums)
}
