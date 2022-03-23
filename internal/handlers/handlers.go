package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Reticent93/Tonic-Massage/internal/config"
	"github.com/Reticent93/Tonic-Massage/internal/driver"
	"github.com/Reticent93/Tonic-Massage/internal/forms"
	"github.com/Reticent93/Tonic-Massage/internal/helpers"
	"github.com/Reticent93/Tonic-Massage/internal/models"
	"github.com/Reticent93/Tonic-Massage/internal/render"
	"github.com/Reticent93/Tonic-Massage/internal/repository"
	"github.com/Reticent93/Tonic-Massage/internal/repository/dbrepo"
	"net/http"
	"strconv"
	"time"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

//NewRepo create a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//send the data to the template
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Massage(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "therapists.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Lymphatic(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "lymphatic.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Swedish(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "swedish.page.tmpl", &models.TemplateData{})
}

func (m *Repository) DeepTissue(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "deep-tissue.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search.avail.page.tmpl", &models.TemplateData{})
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
	var emptyBooking models.Reservation
	data := make(map[string]interface{})
	data["booking"] = emptyBooking

	render.Template(w, r, "booking.page.tmpl", &models.TemplateData{
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

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	layout := "01-02-2006"
	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, err)
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
	}

	therapistId, err := strconv.Atoi(r.Form.Get("therapist_id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	booking := models.Reservation{
		FirstName:   r.Form.Get("first_name"),
		LastName:    r.Form.Get("last_name"),
		Email:       r.Form.Get("email"),
		Phone:       r.Form.Get("phone"),
		StartDate:   startDate,
		EndDate:     endDate,
		TherapistId: therapistId,
	}

	form := forms.New(r.PostForm)
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	form.Required("first_name", "last_name", "email", "phone", "date")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["booking"] = booking

		render.Template(w, r, "booking.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	err = m.DB.InsertReservation(booking)
	if err != nil {
		helpers.ServerError(w, err)
	}

	m.App.Session.Put(r.Context(), "booking", booking)
	http.Redirect(w, r, "/booking-summary", http.StatusSeeOther)

}

func (m *Repository) BookingSummary(w http.ResponseWriter, r *http.Request) {
	booking, ok := m.App.Session.Get(r.Context(), "booking").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("cannot get error from session")
		m.App.Session.Put(r.Context(), "error", "Cannot get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "booking")
	data := make(map[string]interface{})
	data["booking"] = booking

	render.Template(w, r, "booking-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
