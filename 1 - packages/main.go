package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("hello world go")
	auxiliar.Write()
	erro := checkmail.ValidateFormat("dedito182gmail.com")
	fmt.Println(erro)
}
