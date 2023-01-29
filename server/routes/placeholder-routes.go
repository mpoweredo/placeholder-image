package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"placeholder-image/handlers"
)

var PlaceholderRoutes = func(router *mux.Router) {
	router.HandleFunc("/{resolution}", handlers.GetImage).Methods(http.MethodGet)
}
