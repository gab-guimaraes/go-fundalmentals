package main

import "fmt"

func main() {
	//explicit type
	var name string = "scooby-doo"
	//type inference
	name2 := "john wick"
	fmt.Println(name)
	fmt.Println(name2)
	var1, var2 := "v1", "v2"
	fmt.Println(var1, var2)
	const const1 string = "const"
	var1, var2 = var2, var1
}
