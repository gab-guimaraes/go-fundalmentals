package main

import "fmt"

type usuario struct {
	nome    string
	idade   uint8
	address address
}

type address struct {
	address string
	number  uint8
}

func printAddress(a address) {
	fmt.Printf("street: " + a.address)
}

func main() {
	fmt.Println("Struct Files...")
	//how to create struct
	var u usuario
	//{0} default val
	u.idade = 20
	u.nome = "John"
	fmt.Println(u)

	var a address
	a.address = "st 1"
	a.number = 2

	//inference for struct
	u2 := usuario{"Scooby", 2, a}
	fmt.Println(u2)

	u3 := usuario{nome: "wick"}
	fmt.Println(u3)
}
