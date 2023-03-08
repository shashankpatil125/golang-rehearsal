package main

import "fmt"

func main() {
	var x int = 3
	switch {
	case x%2 == 0:
		fmt.Println(x, "is even")
	default:
		fmt.Println("x is odd")
	}

}
