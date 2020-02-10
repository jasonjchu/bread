package routes

import (
	"github.com/go-chi/chi"
	"github.com/jasonjchu/bread/app/handlers/getEmployerHandler"
	"github.com/jasonjchu/bread/app/handlers/getJobsHandler"
)

func InitRoutes(r chi.Router) {
	initJobRoutes(r)
	initEmployerRoutes(r)
}

func initJobRoutes(r chi.Router) {
	r.Route("/jobs", func(r chi.Router) {
		r.Get(getJobsHandler.RouteURL, getJobsHandler.Handler)
	})
}

func initEmployerRoutes(r chi.Router) {
	r.Route("/employers", func(r chi.Router) {
		r.Get(getEmployerHandler.RouteURL, getEmployerHandler.Handler)
	})
}
