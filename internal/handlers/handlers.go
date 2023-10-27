package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SquaredR98/bookings-be/internal/config"
	"github.com/SquaredR98/bookings-be/internal/forms"
	"github.com/SquaredR98/bookings-be/internal/models"
	"github.com/SquaredR98/bookings-be/internal/render"
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
	render.RenderTemplates(w, r, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"
	remoteIp := m.App.Session.GetString(r.Context(), "remoteIp")
	stringMap["remoteIp"] = remoteIp
	render.RenderTemplates(w, r, "general.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) GeneralSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "general.page.tmpl", &models.TemplateData{})
}
func (m *Repository) PremiumSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "premium.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "availability.page.tmpl", &models.TemplateData{})
}
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("startDate")
	end := r.Form.Get("endDate")
	w.Write([]byte(fmt.Sprintf("Start Date: %s, End Date: %s", start, end)))
}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplates(w, r, "reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Something went wrong", err)
	}
	reservation := models.Reservation{
		Fullname: r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Phone:    r.Form.Get("mobile"),
	}

	form := forms.New(r.PostForm)

	// form.Has("name", r)
	form.Required("name", "email")
	form.MinLength("name", 3, r)
	form.IsEmail("email")

	log.Println("React=hed Here")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplates(w, r, "reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}
