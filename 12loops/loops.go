package main

import "fmt"

func main() {

	arr := []int{10, 20, 30, 40}
	for i, value := range arr {
		fmt.Println(i, "th element is : ", value)
	}

}
