package handlers

import (
	"fmt"
	"net/http"

	"github.com/SquaredR98/bookings-be/pkg/config"
	"github.com/SquaredR98/bookings-be/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl")
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}
