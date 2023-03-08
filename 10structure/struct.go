package main

import "fmt"

func main() {

	stud1 := Student{12, "shash", 23}
	stud2 := Student{22, "tejas", 16}
	fmt.Printf("%T\n", stud1)
	fmt.Println(stud1)
	arr := []Student{stud1, stud2}
	fmt.Println(arr)

}

type Student struct {
	Rollno int
	Name   string
	Age    int
}
