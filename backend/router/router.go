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

		sub.Route("/settings", func(sub1.Router) {
			sub1("/", views.GetUserSettings)
			sub1ch("/", views.PatchUserSettings)
		})
	})

	//Org
	r.Route("/org", func(sub chi.Router) {
		sub.Get("/", views.GetOrg)
		sub.Post("/", views.NewOrg)
		sub.Patch("/", views.UpdateOrg)
		sub.Delete("/", views.DeleteOrg)

		sub.Route("/settings", func(sub1.Router) {
			sub1("/", views.GetOrgSettings)
			sub1ch("/", views.PatchOrgSettings)
		})
	})

	//Page
	r.Route("/page", func(sub chi.Router) {
		sub.Get("/", views.GetPage)
		sub.Post("/", views.NewPage)
		sub.Patch("/", views.UpdatePage)
		sub.Delete("/", views.DeletePage)

})
}