package handlers

import (
	"net/http"

	"github.com/JDysiewicz/go-course/pkg/config"
	"github.com/JDysiewicz/go-course/pkg/models"
	"github.com/JDysiewicz/go-course/pkg/render"
)



var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}

