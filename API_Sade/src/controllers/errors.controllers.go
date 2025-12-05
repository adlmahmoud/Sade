package controllers

import (
	"Tp_Api_spotify/models"
	"Tp_Api_spotify/templates"
	"net/http"
)

// ErrorDisplay gère l'affichage de la page d'erreur.
// On récupère le code et le message passés en query string /error?code=404&message=Page%30introuvable
func ErrorDisplay(w http.ResponseWriter, r *http.Request) {
	// Récupération "simple" (string) depuis la query. Exemple: /error?code=404&message=Page%30introuvable
	data := models.Error{
		Code:    r.FormValue("code"),    // "404"
		Message: r.FormValue("message"), // "Page introuvable"
	}

	// On délègue le rendu au helper de templates (buffer + redirection si erreur de rendu)
	templates.RenderTemplate(w, r, "error", data)
}
