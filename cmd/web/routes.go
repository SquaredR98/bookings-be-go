package main

import (
	"net/http"

	"github.com/SquaredR98/bookings-be/internal/config"
	"github.com/SquaredR98/bookings-be/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/general", http.HandlerFunc(handlers.Repo.GeneralSuite))
	mux.Get("/premium", http.HandlerFunc(handlers.Repo.PremiumSuite))
	mux.Get("/availability", http.HandlerFunc(handlers.Repo.Availability))
	mux.Post("/availability", http.HandlerFunc(handlers.Repo.PostAvailability))
	mux.Get("/make-reservation", http.HandlerFunc(handlers.Repo.Reservation))
	mux.Get("/reservation-summary", http.HandlerFunc(handlers.Repo.ReservationSummary))
	mux.Post("/make-reservation", http.HandlerFunc(handlers.Repo.PostReservation))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
