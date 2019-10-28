package main

import (
	"fmt"
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

const port = 8080

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

	fmt.Printf("\033[1mRunning server on port %d\033[0m \n", port)

	if http.ListenAndServe(fmt.Sprintf(":%d", port), r) != nil {
		os.Exit(1)
	}
}
