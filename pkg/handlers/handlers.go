package handlers

import (
	"github.com/Reticent93/Tonic-Massage/internal/config"
	"github.com/Reticent93/Tonic-Massage/internal/models"
	"github.com/Reticent93/Tonic-Massage/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	//send the data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Massage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "therapists.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Lymphatic(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "lymphatic.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Swedish(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "swedish.page.tmpl", &models.TemplateData{})
}

func (m *Repository) DeepTissue(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "deep-tissue.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Booking(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{})
}