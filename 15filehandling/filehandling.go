package main

import (
	"fmt"
	"os"
)

func main() {
	a, e := os.Create("new.txt")
	a.WriteString("kdshklrhfaksdhfk")
	fmt.Printf("%T", a)
	fmt.Println(e)
	os.Mkdir("new", 0223)
	defer a.Close()
}
