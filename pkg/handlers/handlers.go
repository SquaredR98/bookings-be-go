package handlers

import (
	"net/http"

	"github.com/SquaredR98/bookings-be/pkg/config"
	"github.com/SquaredR98/bookings-be/pkg/models"
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
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIp", remoteIp)
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"
	remoteIp := m.App.Session.GetString(r.Context(), "remoteIp")
	stringMap["remoteIp"] = remoteIp
	render.RenderTemplates(w, "general.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
