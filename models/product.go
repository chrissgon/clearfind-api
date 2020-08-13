package models

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/chrissgon/clearfind-api/config"
	"github.com/chrissgon/clearfind-api/utils"
)

/*
########################################
	STRUCTS
########################################
*/

type Product struct {
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Price       struct {
		Real    int `json:"Real"`
		Decimal int `json:"Decimal"`
	} `json:"Price"`
	Path  string `json:"Path"`
	Image string `json:"Image"`
	File
	Quantity int `json:"Quantity"`
}

type Products []Product

/*
########################################
	LEITURA
########################################
*/

// FindProducts - Filtra Produtos
func FindProducts(name string) Products {
	var (
		db            *sql.DB
		err           error
		money         string
		p             Product
		ps            Products
		query         *sql.Rows
		real, decimal int
	)

	// Cria Conexão
	db = config.ConnectDB()

	// Cria Slice
	ps = make(Products, 0)

	// Inicia Query e Trata Erros
	if query, err = db.Query("SELECT * FROM product WHERE product.name LIKE ?", "%"+name+"%"); err != nil {

		// Exibe Erro
		utils.ShowError("FindProducts", 35, err)
	}

	// Encerra Query
	defer query.Close()

	// Inicia Loop
	for query.Next() {

		// Inseri Dados no Struct
		query.Scan(&p.ID, &p.Path, &p.Image, &p.Name, &p.Description, &money, &p.Quantity)

		// Trata Valor
		real, _ = strconv.Atoi(strings.Split(money, ".")[0])
		decimal, _ = strconv.Atoi(strings.Split(money, ".")[1])

		// Inseri Valor no Struct
		p.Price.Real = real
		p.Price.Decimal = decimal

		// Adiciona Struct no Array
		ps = append(ps, p)
	}

	return ps
}

// GetProducts - Pega Produtos
func GetProducts() Products {
	var (
		db            *sql.DB
		err           error
		money         string
		p             Product
		ps            Products
		query         *sql.Rows
		real, decimal int
	)

	// Cria Conexão
	db = config.ConnectDB()

	// Cria Slice
	ps = make(Products, 0)

	if query, err = db.Query("SELECT * FROM product ORDER BY name"); err != nil {

		// Exibe Erro
		utils.ShowError("GetProducts", 106, err)
	}

	// Encerra Query
	defer query.Close()

	// Inicia Loop
	for query.Next() {

		// Inseri Dados no Struct
		query.Scan(&p.ID, &p.Path, &p.Image, &p.Name, &p.Description, &money, &p.Quantity)

		// Trata Valor
		real, _ = strconv.Atoi(strings.Split(money, ".")[0])
		decimal, _ = strconv.Atoi(strings.Split(money, ".")[1])

		// Inseri Valor no Struct
		p.Price.Real = real
		p.Price.Decimal = decimal

		// Adiciona Struct no Array
		ps = append(ps, p)
	}

	return ps
}

// GetProduct - Pega Produto
func GetProduct(ID int) Product {
	var (
		db            *sql.DB
		err           error
		money         string
		p             Product
		real, decimal int
	)

	// Cria Conexão
	db = config.ConnectDB()

	if err = db.QueryRow("SELECT * FROM product WHERE id = ?", ID).Scan(&p.ID, &p.Path, &p.Image, &p.Name, &p.Description, &money, &p.Quantity); err != nil {

		// Exibe Erro
		utils.ShowError("GetProduct", 152, err)
	}

	// Trata Valor
	real, _ = strconv.Atoi(strings.Split(money, ".")[0])
	decimal, _ = strconv.Atoi(strings.Split(money, ".")[1])

	// Inseri Valor no Struct
	p.Price.Real = real
	p.Price.Decimal = decimal

	return p
}

/*
########################################
	CRIAÇÃO
########################################
*/

// CreateProduct - Cria Produto
func CreateProduct(r *http.Request) bool {
	var (
		ID       int
		p        Product
		validate bool
	)

	// Adquire Valores do Formulario
	r.ParseMultipartForm(0)

	// Atribui Valores ao Struct
	p = setFieldsProduct(r)

	// Insere Produto e Trata Erros
	if validate, ID = insertProduct(p); validate == false {

		return false
	}

	// Atribui ID do Produto
	p.ID = ID

	// Cria Arquivo e Trata Erros
	if validate = createFile(p); validate == false {

		return false
	}

	return true
}

func insertProduct(p Product) (bool, int) {
	var (
		db     *sql.DB
		err    error
		ID     int64
		money  string
		stmt   *sql.Stmt
		rows   int64
		result sql.Result
	)

	// Cria Conexão
	db = config.ConnectDB()

	// Inicia Prepare e Trata Erros
	if stmt, err = db.Prepare("INSERT INTO product VALUES(NULL, ?, ?, ?, ?, ?, ?)"); err != nil {

		// Exibe Erro
		utils.ShowError("insertProduct", 212, err)
	}

	money = strconv.Itoa(p.Price.Real) + "." + strconv.Itoa(p.Price.Decimal)

	// Executa Prepare e Trata Erros
	if result, err = stmt.Exec("", "", p.Name, p.Description, money, p.Quantity); err != nil {

		// Exibe Erro
		utils.ShowError("insertProduct", 219, err)
	}

	// Verifica Colunas Afetadas
	if rows, _ = result.RowsAffected(); rows != 0 {

		ID, _ = result.LastInsertId()

		return true, int(ID)
	}

	return false, 0
}

/*
########################################
	EDIÇÃO
########################################
*/

// EditProduct - Edita Produto
func EditProduct(r *http.Request) bool {
	var (
		p        Product
		validate bool
	)

	// Adquire Valores do Formulario
	r.ParseMultipartForm(0)

	// Atribui Valores ao Struct
	p = setFieldsProduct(r)

	// Edita Produto e Trata Erros
	if validate = updateProduct(p); validate == false {

		return false
	}

	// Verifica Existencia de Nova Imagem
	if p.File.Err == nil {

		// Edita Arquivo e Trata Erros
		if validate = createFile(p); validate == false {

			return false
		}
	}

	return true
}

// updateProduct - Atualiza Produto no Banco
func updateProduct(p Product) bool {
	var (
		db    *sql.DB
		err   error
		money string
		stmt  *sql.Stmt
	)

	// Cria Conexão
	db = config.ConnectDB()

	// Inicia Prepare e Trata Erros
	if stmt, err = db.Prepare("UPDATE product SET name = ?, description = ?, price = ?, quantity = ? WHERE id = ?"); err != nil {

		// Exibe Erro
		utils.ShowError("updateProduct", 302, err)

		return false
	}

	money = strconv.Itoa(p.Price.Real) + "." + strconv.Itoa(p.Price.Decimal)

	// Executa Prepare e Trata Erros
	if _, err = stmt.Exec(p.Name, p.Description, money, p.Quantity, p.ID); err != nil {

		// Exibe Erro
		utils.ShowError("updateProduct", 311, err)

		return false
	}

	return true
}

/*
########################################
	EXCLUSÃO
########################################
*/

// DeleteProduct - Deleta Produto
func DeleteProduct(IDs []string) bool {
	var (
		db     *sql.DB
		err    error
		index  int
		ID     string
		params []interface{}
		rows   int64
		result sql.Result
		stmt   *sql.Stmt
		sql    strings.Builder
	)

	// Cria Conexão
	db = config.ConnectDB()

	// Cria Sql Dinamico
	sql.WriteString("DELETE FROM product WHERE id IN(")

	for index, ID = range IDs {
		index++

		// Verifica Ultimo Laco
		if len(IDs) == index {
			sql.WriteString("?")
		} else {
			sql.WriteString("?,")
		}

		// Atribui IDS em Interface
		params = append([]interface{}{ID}, params...)
	}

	sql.WriteString(");")

	// Inicia Prepare e Trata Erros
	if stmt, err = db.Prepare(sql.String()); err != nil {

		// Exibe Erro
		utils.ShowError("DeleteProduct", 348, err)

		return false
	}

	// Executa Prepare e Trata Erros
	if result, err = stmt.Exec(params...); err != nil {

		// Exibe Erro
		utils.ShowError("DeleteProduct", 371, err)

		return false
	}

	// Verifica Colunas Afetadas
	if rows, _ = result.RowsAffected(); rows != 0 {

		return true
	}

	return false
}

/*
########################################
	GENÉRICAS
########################################
*/

// setFieldsProduct - Seta Valores no Struct de Produto
func setFieldsProduct(r *http.Request) Product {
	var (
		money                 string
		p                     Product
		real, decimal, verify int
	)

	// Obtem Valores
	p.ID, _ = strconv.Atoi(r.PostFormValue("id"))
	p.File.File, p.File.Handler, p.File.Err = r.FormFile("imagem")
	p.Name = r.PostFormValue("nome")
	p.Description = r.PostFormValue("descricao")
	p.Quantity, _ = strconv.Atoi(r.PostFormValue("quantidade"))

	// Trata Preço
	money = r.PostFormValue("preco")

	real, _ = strconv.Atoi(strings.Split(money, ",")[0])

	verify = strings.Index(strings.Split(money, ",")[0], ".")

	// Valida Caracteres
	if verify != -1 {

		// Remove Ponto
		real, _ = strconv.Atoi(strings.ReplaceAll(strings.Split(money, ",")[0], ".", ""))
	}

	decimal, _ = strconv.Atoi(strings.Split(money, ",")[1])

	// Insere Valor no Struct
	p.Price.Real = real
	p.Price.Decimal = decimal

	return p
}
