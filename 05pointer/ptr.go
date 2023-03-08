package main

import "fmt"

func main() {
	var num int = 23
	fmt.Println(num)
	var ptr *int = &num
	fmt.Println(*ptr)
}
