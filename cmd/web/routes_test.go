package main

import (
	"testing"

	"github.com/SquaredR98/bookings-be/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoot(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:

	default:
		t.Errorf("Type is not http.Handler, but is %T", v)
	}
}
