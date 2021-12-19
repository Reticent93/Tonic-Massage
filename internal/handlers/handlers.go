package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Reticent93/Tonic-Massage/internal/config"
	"github.com/Reticent93/Tonic-Massage/internal/forms"
	"github.com/Reticent93/Tonic-Massage/internal/helpers"
	"github.com/Reticent93/Tonic-Massage/internal/models"
	"github.com/Reticent93/Tonic-Massage/internal/render"
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
	//send the data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
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

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search.avail.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	therapist := r.Form.Get("therapist")
	book_date := r.Form.Get("book-date")

	w.Write([]byte(fmt.Sprintf("therapist is %s and book-date is %s", therapist, book_date)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//AvailabilityJSON handles request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Booking(w http.ResponseWriter, r *http.Request) {
	var emptyBooking models.Booking
	data := make(map[string]interface{})
	data["booking"] = emptyBooking

	render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

//PostBooking handles the posting of a booking form
func (m *Repository) PostBooking(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	booking := models.Booking{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		Date:      r.Form.Get("date"),
		Time:      r.Form.Get("time"),
	}

	form := forms.New(r.PostForm)
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	form.Required("first_name", "last_name", "email", "phone", "date")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["booking"] = booking

		render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "booking", booking)
	http.Redirect(w, r, "/booking-summary", http.StatusSeeOther)

}

func (m *Repository) BookingSummary(w http.ResponseWriter, r *http.Request) {
	booking, ok := m.App.Session.Get(r.Context(), "booking").(models.Booking)
	if !ok {
		m.App.ErrorLog.Println("cannot get error from session")
		m.App.Session.Put(r.Context(), "error", "Cannot get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "booking")
	data := make(map[string]interface{})
	data["booking"] = booking

	render.RenderTemplate(w, r, "booking-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
