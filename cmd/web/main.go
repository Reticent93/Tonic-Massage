package main

import (
	"encoding/gob"
	"fmt"
	"github.com/Reticent93/Tonic-Massage/internal/config"
	"github.com/Reticent93/Tonic-Massage/internal/handlers"
	"github.com/Reticent93/Tonic-Massage/internal/models"
	"github.com/Reticent93/Tonic-Massage/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting app on port %s", port))
	//_ = http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	//what I put in the session
	gob.Register(models.Booking{})

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	return nil
}
