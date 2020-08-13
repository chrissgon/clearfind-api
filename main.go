package main

import (
	"net/http"

	"github.com/chrissgon/clearfind-api/middlewares"
	"github.com/chrissgon/clearfind-api/routes"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	var (
		n      *negroni.Negroni
		router *mux.Router
	)

	// Iniciliaza Servidor
	n = negroni.Classic()

	// Define Content-Type
	n.Use(middlewares.ApplicationJSON())

	// Habilita Cors
	n.Use(middlewares.EnableCors())

	// Cria Router
	router = mux.NewRouter()
	n.UseHandler(router)

	router.Handle("/docs/{image}", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))

	// Obtem Caminhos
	routes.Routes(router)

	// Inicializa Porta
	n.Run(":3333")
}
