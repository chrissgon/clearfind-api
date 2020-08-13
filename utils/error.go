package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ShowError(method string, line int, err error) {

	// Exibe Erro
	fmt.Println("------------------------")
	fmt.Println("Erro no metodo " + method + ", linha " + strconv.Itoa(line) + ".")
	fmt.Println("------------------------")
	fmt.Println(err.Error())
	os.Exit(1)
}
