package routes

import (
	"github.com/go-chi/chi"
	"github.com/jasonjchu/bread/app/handlers/candidateDislikesJobHandler"
	"github.com/jasonjchu/bread/app/handlers/candidateLikesJobHandler"
	"github.com/jasonjchu/bread/app/handlers/candidateLoginHandler"
	"github.com/jasonjchu/bread/app/handlers/candidateRegisterHandler"
	"github.com/jasonjchu/bread/app/handlers/employerLoginHandler"
	"github.com/jasonjchu/bread/app/handlers/employerRegisterHandler"
	"github.com/jasonjchu/bread/app/handlers/getCandidatesForJobHandler"
	"github.com/jasonjchu/bread/app/handlers/getCompaniesHandler"
	"github.com/jasonjchu/bread/app/handlers/getEmployerHandler"
	"github.com/jasonjchu/bread/app/handlers/getEmployerMatchHandler"
	"github.com/jasonjchu/bread/app/handlers/getJobsForCandidatesHandler"
	"github.com/jasonjchu/bread/app/handlers/getJobsForEmployerHandler"
	"github.com/jasonjchu/bread/app/handlers/getJobsHandler"
)

func InitRoutes(r chi.Router) {
	initJobRoutes(r)
	initEmployerRoutes(r)
	initCandidateRoutes(r)
	initCompanyRoutes(r)
}

func initJobRoutes(r chi.Router) {
	r.Route("/jobs", func(r chi.Router) {
		r.Get(getJobsHandler.RouteURL, getJobsHandler.Handler)
	})
}

func initEmployerRoutes(r chi.Router) {
	r.Route("/employers", func(r chi.Router) {
		r.Get(getEmployerHandler.RouteURL, getEmployerHandler.Handler)
		r.Post(employerLoginHandler.RouteURL, employerLoginHandler.Handler)
		r.Post(employerRegisterHandler.RouteURL, employerRegisterHandler.Handler)
		r.Route("/candidates-for-job", func(r chi.Router) {
			r.Get(getCandidatesForJobHandler.RouteURL, getCandidatesForJobHandler.Handler)
		})
		r.Get(getJobsForEmployerHandler.RouteURL, getJobsForEmployerHandler.Handler)
		r.Get(getEmployerMatchHandler.RouteURL, getEmployerMatchHandler.Handler)
	})
}

func initCandidateRoutes(r chi.Router) {
	r.Route("/candidates", func(r chi.Router) {
		r.Post(candidateRegisterHandler.RouteURL, candidateRegisterHandler.Handler)
		r.Post(candidateLoginHandler.RouteURL, candidateLoginHandler.Handler)

		r.Route("/jobs", func(r chi.Router) {
			// GET JOBS
			r.Get(getJobsForCandidatesHandler.RouteURL, getJobsForCandidatesHandler.Handler)

			// LIKE AND DISLIKE JOBS
			r.Route("/{job_id}", func(r chi.Router) {
				r.Post(candidateLikesJobHandler.RouteURL, candidateLikesJobHandler.Handler)
				r.Post(candidateDislikesJobHandler.RouteURL, candidateDislikesJobHandler.Handler)
			})
		})
	})
}

func initCompanyRoutes(r chi.Router) {
	r.Route("/companies", func(r chi.Router) {
		r.Get(getCompaniesHandler.RouteURL, getCompaniesHandler.Handler)
	})
}
