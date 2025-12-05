package routers

import (
	"Tp_Api_spotify/controllers"
	"net/http"
)

func TracksRouter(router *http.ServeMux) {
	router.HandleFunc("/best_track", controllers.DisplayTrack)
}
