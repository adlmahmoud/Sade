package controllers

import (
	"Tp_Api_spotify/helpers"
	"Tp_Api_spotify/services"
	"Tp_Api_spotify/templates"
	"fmt"
	"net/http"
)

func DisplayArtiste(w http.ResponseWriter, r *http.Request) {
	const id_Sade = "47zz7sob9NUcODy0BTDvKx"
	artiste, status, err := services.SearchArtisteSrvices(id_Sade)
	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération de l'artiste")
		fmt.Println(err)
		return
	}
	templates.RenderTemplate(w, r, "artiste", artiste)
}
