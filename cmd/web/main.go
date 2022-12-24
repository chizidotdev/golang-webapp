package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chizidotdev/golang-webapp/pkg/config"
	"github.com/chizidotdev/golang-webapp/pkg/handlers"
	"github.com/chizidotdev/golang-webapp/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTempltes(&app)

	fmt.Printf("App started on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
