package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KoksalBerkay/go-review/go-webapp-sample/pkg/config"
	"github.com/KoksalBerkay/go-review/go-webapp-sample/pkg/handlers"
	"github.com/KoksalBerkay/go-review/go-webapp-sample/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080" //const value cannot be changed
var app config.AppConfig
var session *scs.SessionManager

// main is the entry point for the program
func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
