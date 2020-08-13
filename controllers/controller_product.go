package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/chrissgon/clearfind-api/models"
	"github.com/gorilla/mux"
)

// FindProducts - Filtra Produtos
func FindProducts(rw http.ResponseWriter, r *http.Request) {
	var (
		result   []byte
		name     string
		products models.Products
		vars     map[string]string
	)

	// Adquire Variaveis
	vars = mux.Vars(r)

	// Recupera Name
	name, _ = url.QueryUnescape(vars["name"])

	// Verifica Valor Filtrado
	if name != "null" {

		// Filtra Produtos
		products = models.FindProducts(name)
	} else {

		// Pega Produtos
		products = models.GetProducts()
	}

	// Converte para Json
	result, _ = json.Marshal(&products)

	rw.Write(result)
}

// GetProducts - Pega Produtos
func GetProducts(rw http.ResponseWriter, r *http.Request) {
	var (
		products models.Products
		result   []byte
	)

	// Pega Produtos
	products = models.GetProducts()

	result, _ = json.Marshal(&products)

	rw.Write(result)
}

// GetProduct - Pega Produto
func GetProduct(rw http.ResponseWriter, r *http.Request) {
	var (
		result  []byte
		ID      int
		product models.Product
		vars    map[string]string
	)

	// Adquire Variaveis
	vars = mux.Vars(r)

	// Recupera ID
	ID, _ = strconv.Atoi(vars["id"])

	// Pega Produto
	product = models.GetProduct(ID)

	// Converte para Json
	result, _ = json.Marshal(&product)

	rw.Write(result)
}

// CreateProduct - Cria Produto
func CreateProduct(rw http.ResponseWriter, r *http.Request) {
	var (
		validate bool
		result   []byte
	)

	// Pega Resultado
	validate = models.CreateProduct(r)

	// Converte para Json
	result, _ = json.Marshal(&validate)

	rw.Write(result)
}

// EditProduct - Edita Produto
func EditProduct(rw http.ResponseWriter, r *http.Request) {
	var (
		validate bool
		result   []byte
	)

	// Pega Resultado
	validate = models.EditProduct(r)

	// Converte para Json
	result, _ = json.Marshal(&validate)

	rw.Write(result)
}

// DeleteProduct - Deleta Produto
func DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	var (
		result   []byte
		ID       string
		IDs      []string
		validate bool
		vars     map[string]string
	)

	// Adquire Variaveis
	vars = mux.Vars(r)

	// Recupera ID
	ID = vars["id"]

	// Desmenbra IDS
	IDs = strings.Split(ID, ",")

	// Deleta Produto
	validate = models.DeleteProduct(IDs)

	// Converte para Json
	result, _ = json.Marshal(&validate)

	rw.Write(result)
}
