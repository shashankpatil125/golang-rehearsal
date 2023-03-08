package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = make([]int, 4)
	arr[0] = 675
	arr[1] = 233
	arr[2] = 345
	arr[3] = 132
	fmt.Println(arr)
	arr = append(arr, 232, 65, 234)
	//here we are appending some element to that array
	// after appending it will realocate in a memory
	// thats why we are not get a out of bound error
	fmt.Println(arr)
	fmt.Println(arr[0:3])
	sort.Ints(arr)
	fmt.Println(arr)

	// slice is not a array it will just refer to a subarray
	// if we will chenging something at slice array then it will reflect to main array

	rollno := [5]int{43, 45, 6, 7, 4}
	fmt.Println(rollno)
	newslice, s2 := rollno[:2], rollno[3:]
	fmt.Println(newslice)
	fmt.Println(s2)
	newslice = append(newslice, 23, 45, 645, 345, 45)
	fmt.Println(newslice)

	a := []int{23, 23, 23, 23}
	b := []int{34, 34, 2, 3, 34, 34}
	fmt.Println(a, b)
	ab := append(a, b...)
	fmt.Println(ab)
	fmt.Printf("%T", ab)
	fmt.Printf("%T\n", rollno)

	// if there are we give a size for the array then it will make a array
	//otherwise it will the array

	// delete a element from an array
	data := [5]int{23, 34, 45, 5, 6}
	fmt.Println(data)
	//delete 45
	res := data[:2]
	res2 := data[3:]
	ans := append(res, res2...)
	fmt.Println(ans)

}
