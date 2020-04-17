package router

import (
	"github.com/gorilla/mux"
	"github.com/lusinx/pressx/views"
)

// Route Main Router
func Route(r *mux.Router) {
	// User routes

	r.HandleFunc("/user/{user}", views.GetUser).Methods("GET") // Public
	r.HandleFunc("/user", views.NewUser).Methods("POST")
	r.HandleFunc("/user", views.UpdateUser).Methods("PUT")
	r.HandleFunc("/user", views.DeleteUser).Methods("DELETE")

	// user settings routes
	r.HandleFunc("/user/settings", views.GetUserSettings).Methods("GET")
	r.HandleFunc("/user/settings", views.UpdateUserSettings).Methods("PUT")

	// Org routes
	r.HandleFunc("/org/{org}", views.GetOrg).Methods("GET") // Public
	r.HandleFunc("/org", views.NewOrg).Methods("GET")
	r.HandleFunc("/org", views.UpdateOrg).Methods("PUT")
	r.HandleFunc("/user", views.DeleteOrg).Methods("GET")

	// org settings routes
	r.HandleFunc("/org/settings", views.GetOrgSettings).Methods("GET")
	r.HandleFunc("/org/settings", views.UpdateOrgSettings).Methods("PUT")

	// Post routes
	r.HandleFunc("/page/{page}", views.GetPage).Methods("GET") // Public
	r.HandleFunc("/page", views.NewPage).Methods("POST")
	r.HandleFunc("/page/{page}", views.UpdatePage).Methods("PUT")
	r.HandleFunc("/page/{page}", views.DeletePage).Methods("POST")
}

// r.Group(func(r chi.Router) {
// 	route := ""
// 	r.Use(
// 		jwtauth.Verifier(session.TokenAuth),
// 		session.ModAuthenticator(route),
// 	)
// 	// User
// 	// r.Route("/user", func(sub mux.Router) {
// 	// 	route = "/user"
// 	// 	sub.Get("/", views.GetUser)
// 	// 	sub.Post("/", views.NewUser)
// 	// 	sub.Patch("/", views.UpdateUser)
// 	// 	sub.Delete("/", views.DeleteUser)

// 	// 	sub.Route("/settings", func(sub chi.Router) {
// 	// 		sub.Get("/", views.GetUserSettings)
// 	// 		sub.Patch("/", views.PatchUserSettings)
// 	// 	})
// 	// })

// 	r.Route("/org", func(sub chi.Router) {
// 		route = "/org"
// 		sub.Post("/", views.NewOrg)
// 		sub.Patch("/", views.UpdateOrg)
// 		sub.Delete("/", views.DeleteOrg)

// 		sub.Route("/settings", func(sub chi.Router) {
// 			sub.Get("/", views.GetOrgSettings)
// 			sub.Patch("/", views.PatchOrgSettings)
// 		})
// 	})

// 	//Page
// 	r.Route("/page", func(sub chi.Router) {
// 		route = "/page"
// 		sub.Post("/", views.NewPage)
// 		sub.Patch("/", views.UpdatePage)
// 		sub.Delete("/", views.DeletePage)
// 	})

// })

// 	// Public routes
// 	r.Group(func(r chi.Router) {
// 		// User
// 		r.Route("/user/{user}", func(sub chi.Router) {
// 			sub.Get("/", views.GetUser)
// 		})

// 		// Org
// 		r.Route("/org/{org}", func(sub chi.Router) {
// 			sub.Get("/", views.GetOrg)
// 		})

// 		// Page
// 		r.Route("/page/{page}", func(sub chi.Router) {
// 			sub.Get("/", views.GetPage)
// 		})
// 	})

// }
