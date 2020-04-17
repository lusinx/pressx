package router

import (
	"github.com/gorilla/mux"
	"github.com/lusinx/pressx/middleware"
	"github.com/lusinx/pressx/views"
)

// Route Main Router
func Route(r *mux.Router) {
	// Level 255 - Pleb number
	// authed are routes that require base authentication
	authed := r.NewRoute().Subrouter()
	authed.Use(middleware.AuthUser(255))

	// User routes
	r.HandleFunc("/user/{user}", views.GetUser).Methods("GET") // Public
<<<<<<< HEAD
	authed.HandleFunc("/user", views.NewUser).Methods("POST")
	authed.HandleFunc("/user", views.UpdateUser).Methods("PATCH")
	authed.HandleFunc("/user", views.DeleteUser).Methods("DELETE")

	// User settings routes
	authed.HandleFunc("/user/settings", views.GetUserSettings).Methods("GET")
	authed.HandleFunc("/user/settings", views.UpdateUserSettings).Methods("PATCH")

	// Org routes
	r.HandleFunc("/org/{org}", views.GetOrg).Methods("GET") // Public
	authed.HandleFunc("/org", views.NewOrg).Methods("GET")
	authed.HandleFunc("/org", views.UpdateOrg).Methods("GET")
	authed.HandleFunc("/user", views.DeleteOrg).Methods("GET")

	// org settings routes
	authed.HandleFunc("/org/settings", views.GetOrgSettings).Methods("GET")
	authed.HandleFunc("/org/settings", views.UpdateOrgSettings).Methods("GET")

	// Post routes
	r.HandleFunc("/page/{page}", views.GetPage).Methods("GET") // Public
	authed.HandleFunc("/page", views.NewPage).Methods("POST")
	authed.HandleFunc("/page/{page}", views.UpdatePage).Methods("POST")
	authed.HandleFunc("/page/{page}", views.DeletePage).Methods("POST")
=======
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
>>>>>>> 160392fcfad38e1c8553933dd506e07188256c63
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
