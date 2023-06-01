package handlers

import (
	"net/http"

	"github.com/salimmia/bookings/pkg/config"
	"github.com/salimmia/bookings/pkg/models"
	"github.com/salimmia/bookings/pkg/render"
)

// Repo is the repository by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the Handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the Home page Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_id", remoteIp)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// build some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_id")

	stringMap["remote_id"] = remoteIp

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
