package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SquaredR98/bookings-be/pkg/config"
	"github.com/SquaredR98/bookings-be/pkg/handlers"
	"github.com/SquaredR98/bookings-be/pkg/render"
)

const PORT = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateFromCache()

	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Printf("Server running on PORT: %s", PORT)

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
