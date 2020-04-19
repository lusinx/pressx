package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	mid "github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lusinx/pressx/config"
	"github.com/lusinx/pressx/models"
	"github.com/lusinx/pressx/router"
)

// Rule of thumb: Strict MVC is overrated.
// Make a functional router, then add functionality to slug patterns.
// Benefits: Less convoluted and easier to maintain.

// Config for Chi Middlewares
type Config struct {
	Timeout time.Duration // In seconds
}

const port = 80

func main() {
	r := mux.NewRouter().StrictSlash(false)

	// Set default config
	timeConf := Config{
		Timeout: 60,
	}

	// Middlewares the default stack shall use
	r.Use(
		mid.RequestID,
		mid.RealIP,
		mid.Logger,
		mid.Recoverer,
		mid.Timeout(timeConf.Timeout*time.Second),
	)
	config.Init()
	models.DBMigrate()
	// Route to views
	router.Route(r)

	fmt.Printf("\033[1mRunning server on port %d\033[0m \n", port)

	if http.ListenAndServe(fmt.Sprintf(":%d", port), r) != nil {
		os.Exit(1)
	}
}
