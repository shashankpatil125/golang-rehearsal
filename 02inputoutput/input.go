//  in other launguages there are try and catch in go it is absent
// instade of try and catch we use the "," and "_"

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("That programming is taking input ")

	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("enter the input: ")
	variable, _ := scanner.ReadString('\n')
	fmt.Println(variable)

	fmt.Println("another way to take input: ")
	var usingScanf int
	fmt.Scanln(&usingScanf)
	fmt.Println(usingScanf)

}
