package main

import "fmt"

func add(a, b int) int { //single value return function
	return a + b
}

func calculator(a, b int) (string, int, string, int, string, int, string, int) { //multiple values return function

	return "\n add : ", a + b, "\n sub : ", a - b, "\n mul : ", a * b, "\n div : ", a / b
}

func main() {
	//single value return function
	a := 10
	b := 20
	fmt.Println(add(a, b))
	fmt.Println(calculator(a, b))

	//multiple values return function
}
