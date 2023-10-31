package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SquaredR98/bookings-be/internal/config"
	"github.com/SquaredR98/bookings-be/internal/handlers"
	"github.com/SquaredR98/bookings-be/internal/helpers"
	"github.com/SquaredR98/bookings-be/internal/models"
	"github.com/SquaredR98/bookings-be/internal/render"
	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"

var app config.AppConfig

var session *scs.SessionManager

var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := runServer()

	if err != nil {
		log.Fatal("Error running server", err)
	}

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

func runServer() error {
	gob.Register(models.Reservation{})

	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateFromCache()

	if err != nil {
		log.Fatal("Cannot create template cache", err)
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)
	helpers.NewHelper(&app)

	return nil
}
