package main

import (
	"github.com/Reticent93/Tonic-Massage/internal/config"
	"github.com/Reticent93/Tonic-Massage/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/massage-therapists", handlers.Repo.Massage)
	mux.Get("/lymphatic", handlers.Repo.Lymphatic)
	mux.Get("/swedish", handlers.Repo.Swedish)
	mux.Get("/deep-tissue", handlers.Repo.DeepTissue)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/booking", handlers.Repo.Booking)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
