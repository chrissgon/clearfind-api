package models

import (
	"database/sql"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/chrissgon/clearfind-api/config"
	"github.com/chrissgon/clearfind-api/utils"
	"github.com/lithammer/shortuuid/v3"
)

/*
########################################
	STRUCTS
########################################
*/

type File struct {
	File    multipart.File        `json:"File"`
	Handler *multipart.FileHeader `json:"Handler"`
	Err     error                 `json:"Err"`
}

type Files []File

/*
########################################
	LEITURA
########################################
*/

func getPath(ID int) string {
	var (
		db   *sql.DB
		err  error
		path string
	)

	// Cria Conexão
	db = config.ConnectDB()

	// Inicia Query e Trata Erros
	if err = db.QueryRow("SELECT path FROM product WHERE id = ?", ID).Scan(&path); err != nil {

		// Exibe Erro
		utils.ShowError("clearPath", 64, err)
	}

	return path
}

/*
########################################
	CRIAÇÃO
########################################
*/

// createFile - Cria Arquivo
func createFile(p Product) bool {
	var (
		validate bool
	)

	// Verifica Existencia do Produto
	if p.ID != 0 {

		// Limpa Imagem Existente
		deleteFile(p)
	}

	// Salva Arquivo no Diretorio
	p.Path, p.Image = uploadFile(p)

	// Atualiza Path do Produto
	validate = updatePath(p)

	return validate
}

// uploadFile - Realiza Upload do Arquivo
func uploadFile(p Product) (string, string) {
	var (
		extension string
		file      *os.File
		hash      string
		name      string
		path      string
	)

	// Adquire Nome Original e Extensao
	extension = "." + strings.Split(p.Handler.Filename, ".")[1]

	name = strings.Split(p.Handler.Filename, ".")[0] + extension

	// Cria Nome Substituto
	hash = shortuuid.New()

	// Caminho do Diretorio
	path = hash + extension

	// Cria Arquivo no Diretorio
	file, _ = os.OpenFile("docs/"+path, os.O_WRONLY|os.O_CREATE, 0666)

	// Encerra File
	defer file.Close()

	// Copia Arquivo no Diretorio
	io.Copy(file, p.File.File)

	return path, name
}

/*
########################################
	EDICAO
########################################
*/

// updatePath - Atualiza Caminho no Banco
func updatePath(p Product) bool {
	var (
		db     *sql.DB
		err    error
		result sql.Result
		rows   int64
		stmt   *sql.Stmt
	)

	// Cria Conexão
	db = config.ConnectDB()

	// Inicia Prepare e Trata Erros
	if stmt, err = db.Prepare("UPDATE product SET path = ?, image = ? WHERE id = ?"); err != nil {

		// Exibe Erro
		utils.ShowError("updatePath", 136, err)
	}

	// Executa Prepare e Trata Erros
	if result, err = stmt.Exec(p.Path, p.Image, p.ID); err != nil {

		// Exibe Erro
		utils.ShowError("updatePath", 143, err)
	}

	// Verifica Linhas Afetadas
	if rows, _ = result.RowsAffected(); rows != 0 {

		return true
	}

	return false
}

/*
########################################
	EXCLUSAO
########################################
*/

// deleteFile - Deleta Arquivo
func deleteFile(p Product) {
	var (
		path string
	)

	// Adquire Path
	path = getPath(p.ID)

	// Remove Path
	removePath(p.ID)

	// Exclui Arquivo
	removeFile(path)

}

// removePath - Remove Arquivo no Banco
func removePath(ID int) {
	var (
		db   *sql.DB
		err  error
		stmt *sql.Stmt
	)

	// Cria Conexão
	db = config.ConnectDB()

	// Inicia Prepare e Trata Erros
	if stmt, err = db.Prepare("UPDATE product SET path = '' WHERE id = ?"); err != nil {

		// Exibe Erro
		utils.ShowError("clearPath", 64, err)
	}

	// Executa Prepare e Trata Erros
	if _, err = stmt.Exec(ID); err != nil {

		// Exibe Erro
		utils.ShowError("clearPath", 71, err)
	}
}

// removeFile - Exclui Arquivo no Diretorio
func removeFile(path string) {

	// Verifica Path
	if path != "" {

		// Remove Arquivo do Diretorio
		os.Remove("docs/" + path)
	}
}
