package main

import (
	"log"
	"net/http"
	"time"

	"github.com/JDysiewicz/go-course/pkg/config"
	"github.com/JDysiewicz/go-course/pkg/handlers"
	"github.com/JDysiewicz/go-course/pkg/render"
	"github.com/alexedwards/scs/v2"
)

// Divide page handler

var app config.AppConfig
var session *scs.SessionManager

func main() {
	

	app.InProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd
	app.Session = session


	tc, err := render.CreateTemplateCache()
	if err !=nil {
		log.Fatal("Cannot get template cache")
	}

	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server {
		Addr: ":8080",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}