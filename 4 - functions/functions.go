package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathcalcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	result := sum(10, 20)
	fmt.Println(result)

	var f = func() {
		fmt.Println("func f")
	}

	var f2 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f()
	f2("test")
	//use _ underscore when you don't need to use the var
	result1, _ := mathcalcs(2, 3)
	fmt.Println(result1)

}
