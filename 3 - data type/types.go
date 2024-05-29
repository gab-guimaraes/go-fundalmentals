package main

import (
	"errors"
	"fmt"
)

func main() {
	//Integer numbers
	//int8, int16, int32, int64
	//int
	var number int32 = 100000000
	fmt.Println(number)

	var number2 uint32 = 200 //without signal
	fmt.Println(number2)

	//float32, float64
	var floatNumber float32 = 123.32
	fmt.Println(floatNumber)

	char := 'B'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	//var boolean bool = true

	//default is <nil>
	var erro error = errors.New("internal error")
	fmt.Println(erro)
}
