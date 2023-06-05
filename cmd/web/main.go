package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/salimmia/bookings/internal/config"
	"github.com/salimmia/bookings/internal/handlers"
	"github.com/salimmia/bookings/internal/render"
	// "os"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// change this to true when in production
	app.InProduction = false
	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// initializing Template cache
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// store template cache to app config
	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/home", handlers.Repo.Home)

	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("%v\n", fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()

	log.Fatal(err)
}
