package main

import "fmt"

func main() {
	pair := make(map[int]string, 5)
	pair[23] = "twentythree" //add pairs
	pair[0] = "zero"         //add pairs
	pair[1] = "one"          //add pairs
	pair[3] = "three"        //add pairs
	pair[2] = "two"          //add pairs
	fmt.Println(pair)
	delete(pair, 2) //delete(mapname, value we want to remove)
	fmt.Println(pair)
	fmt.Printf("%T\n", pair)

	for i, j := range pair {
		fmt.Println(i, j)
	}

}
