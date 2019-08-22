package router

import (
	"github.com/go-chi/chi"
	"github.com/lusinx/pressx/views"
)

func Route(r *chi.Mux) {
	// User
	r.Route("/user", func(sub chi.Router) {
		sub.Get("/", views.GetUser)
		sub.Post("/", views.NewUser)
		sub.Patch("/", views.UpdateUser)
		sub.Delete("/", views.DeleteUser)
	})
}
