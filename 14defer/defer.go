package main

import "fmt"

func check() int {
	defer fmt.Println("one")
	fmt.Println("two")
	return 3
}

func main() {
	fmt.Println(check())
}

//it is used such cases as
// 1 resousces cleanup
// 2 close file
// 3 debugging or testing on database
// 4 panic recovery!
