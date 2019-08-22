package main

import (
	"time"

	"github.com/go-chi/chi"
	mid "github.com/go-chi/chi/middleware"
)

// Rule of thumb: Strict MVC is overrated.
// Make a fucntional router, then add functionality to slug patterns.
// Benefits: Less convoluted and easier to maintain.

type Config struct {
	Timeout time.Duration // In seconds
}

func main() {
	router := chi.NewRouter()

	// Set default config
	config := Config{
		Timeout: 60,
	}

	// Middlewares the default stack shall use
	router.Use(
		mid.RequestID,
		mid.RealIP,
		mid.Logger,
		mid.Recoverer,
		mid.Timeout(config.Timeout*time.Second),
	)

}
