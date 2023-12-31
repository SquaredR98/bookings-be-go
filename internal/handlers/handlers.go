package handlers

import (
	"fmt"
	"net/http"

	"github.com/SquaredR98/bookings-be/internal/config"
	"github.com/SquaredR98/bookings-be/internal/forms"
	"github.com/SquaredR98/bookings-be/internal/helpers"
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
	render.RenderTemplates(w, r, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "about.page.tmpl", &models.TemplateData{})
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
		helpers.ServerError(w, err)
	}
	reservation := models.Reservation{
		Fullname: r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Phone:    r.Form.Get("mobile"),
	}

	form := forms.New(r.PostForm)

	// form.Has("name", r)
	form.Required("name", "email")
	form.MinLength("name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplates(w, r, "reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Can't get error from seesion")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session.")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplates(w, r, "reservation-summary.page.tmpl", &models.TemplateData{Data: data})
}
