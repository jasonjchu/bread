package routes

import (
	"github.com/go-chi/chi"
)

func initJobRoutes(r chi.Router) {
	
}

func initRoutes(r chi.Router) {
	initJobRoutes(r)
}
