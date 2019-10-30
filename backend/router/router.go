package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/lusinx/pressx/session"
	"github.com/lusinx/pressx/views"
)

func Route(r *chi.Mux) {
	// Private routes
	r.Group(func(r chi.Router) {
		r.Use(
			jwtauth.Verifier(session.TokenAuth),
			session.Authenticator,
		)
		// User
		r.Route("/user", func(sub chi.Router) {
			sub.Get("/", views.GetUser)
			sub.Post("/", views.NewUser)
			sub.Patch("/", views.UpdateUser)
			sub.Delete("/", views.DeleteUser)

			sub.Route("/settings", func(sub chi.Router) {
				sub.Get("/", views.GetUserSettings)
				sub.Patch("/", views.PatchUserSettings)
			})
		})

		r.Route("/org", func(sub chi.Router) {
			sub.Post("/", views.NewOrg)
			sub.Patch("/", views.UpdateOrg)
			sub.Delete("/", views.DeleteOrg)

			sub.Route("/settings", func(sub chi.Router) {
				sub.Get("/", views.GetOrgSettings)
				sub.Patch("/", views.PatchOrgSettings)
			})
		})

		//Page
		r.Route("/page", func(sub chi.Router) {
			sub.Post("/", views.NewPage)
			sub.Patch("/", views.UpdatePage)
			sub.Delete("/", views.DeletePage)
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		// User
		r.Route("/user/{user}", func(sub chi.Router) {
			sub.Get("/", views.GetUser)
		})

		// Org
		r.Route("/org/{org}", func(sub chi.Router) {
			sub.Get("/", views.GetOrg)
		})

		// Page
		r.Route("/page/{page}", func(sub chi.Router) {
			sub.Get("/", views.GetPage)
		})
	})

}
