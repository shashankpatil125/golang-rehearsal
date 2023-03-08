// here i am converting the string to integer
// we will use it when we take a data from query/url

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("this is the type casting string to integer ")
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("enter the input")
	input, _ := scanner.ReadString('\n')
	fmt.Println(input)
	n, e := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println(n+34, e)
}
