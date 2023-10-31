package render

import (
	"net/http"
	"testing"

	"github.com/SquaredR98/bookings-be/internal/models"
)

func TestAddDefautlData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()

	if err != nil {
		t.Fatal(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("Failed flash verification")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"

	tc, err := CreateTemplateFromCache()

	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplates(&ww, r, "home.page.tmpl", &models.TemplateData{})

	if err != nil {
		t.Error("Error writing template to browser")
	}

	err = RenderTemplates(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})

	if err == nil {
		t.Error("Rendered template that does not exist")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))

	r = r.WithContext(ctx)
	return r, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplate(app)
}

func TestCreateTemplateFromCache(t *testing.T) {
	pathToTemplates = "./../../templates"

	_, err := CreateTemplateFromCache()

	if err != nil {
		t.Error(err)
	}
}
