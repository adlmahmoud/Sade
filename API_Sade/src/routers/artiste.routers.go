package routers

import (
	"Tp_Api_spotify/controllers"
	"net/http"
)

func ArtisteRouter(router *http.ServeMux) {
	router.HandleFunc("/", controllers.DisplayArtiste)
	router.HandleFunc("/Sade", controllers.DisplayArtiste)
}
