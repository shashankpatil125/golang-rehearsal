package main

import "fmt"

var TT int = 200

func main() {
	fmt.Println("hello")

	var n int = 10
	fmt.Println(n)
	n += 10
	fmt.Println(n)

	var s string = "that is string"
	fmt.Println(s)
	fmt.Printf("the type is %T \n", s)

	var f float32 = 12.212
	fmt.Println(f)

	var unsineint uint8 = 255
	fmt.Println(unsineint)

	var T uint = 25534
	fmt.Println(T)
	fmt.Println(TT)

	random := true
	random = true
	fmt.Println(random)
	fmt.Printf("%T \n", random)
}
