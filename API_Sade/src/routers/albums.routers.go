package routers

import (
	"Tp_Api_spotify/controllers"
	"net/http"
)

func AlbumsRouter(router *http.ServeMux) {
	router.HandleFunc("/albums", controllers.DisplayListAlbums)
}
