package main

import "fmt"

func main() {
	var x interface{}
	x = "This is a string."

	if value, ok := x.(string); ok {
		fmt.Println("Var is string, value = ", value)
	} else {
		fmt.Println("Var is not string.")
	}

	ReadAndWrite()
}
