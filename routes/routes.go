package routes

import (
	"github.com/chrissgon/clearfind-api/controllers"
	"github.com/gorilla/mux"
)

// Routes - Rotas do Sistema
var Routes = func(router *mux.Router) {

	// Rotas Produtos
	router.HandleFunc("/products/find/{name}", controllers.FindProducts).Methods("GET")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/show/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/product/create", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/edit", controllers.EditProduct).Methods("PUT")
	router.HandleFunc("/product/delete/{id}", controllers.DeleteProduct).Methods("DELETE")
}
