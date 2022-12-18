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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("App started on port %s", portNumber))

	http.ListenAndServe(portNumber, nil)
}
