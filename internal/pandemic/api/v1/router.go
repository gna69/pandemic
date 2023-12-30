package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const APIPattern = "/api/v1"

func NewRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/map", GetMapHandler())
	return router
}
