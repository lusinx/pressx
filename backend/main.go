package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	mid "github.com/go-chi/chi/middleware"
	"github.com/lusinx/pressx/router"
)

// Rule of thumb: Strict MVC is overrated.
// Make a functional router, then add functionality to slug patterns.
// Benefits: Less convoluted and easier to maintain.

type Config struct {
	Timeout time.Duration // In seconds
}

func main() {
	r := chi.NewRouter()

	// Set default config
	config := Config{
		Timeout: 60,
	}

	// Middlewares the default stack shall use
	r.Use(
		mid.RequestID,
		mid.RealIP,
		mid.Logger,
		mid.Recoverer,
		mid.Timeout(config.Timeout*time.Second),
	)
	// Route to views
	router.Route(r)

	if http.ListenAndServe(":8080", r) != nil {
		os.Exit(1)
	}
}
