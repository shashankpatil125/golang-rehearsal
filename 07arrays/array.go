package main

import "fmt"

func main() {
	var arr []int
	fmt.Println("length", len(arr))

	// initilized array
	var arr1 = [2]int{1, 2}
	fmt.Println(arr1)

	for i := range arr1 {
		fmt.Println(i)
	}

}
