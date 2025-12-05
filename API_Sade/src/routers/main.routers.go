package routers

import (
	"net/http"
)

func MainRouter() *http.ServeMux {
	mainRouter := http.NewServeMux()

	errorRouter(mainRouter)
	ArtisteRouter(mainRouter)
	AlbumsRouter(mainRouter)
	TracksRouter(mainRouter)

	fileServer := http.FileServer(http.Dir("./../../assets"))
	mainRouter.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mainRouter
}
