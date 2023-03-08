package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Date())
	fmt.Println(t.Day())
	fmt.Println(t.Year())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())

	// date := strconv.Itoa(t.Day()) + "-" + strconv.Itoa(t.Year())
	// fmt.Println(date)
	d := time.Now()
	todaysDate := d.Format("2006-01-02")
	fmt.Println(todaysDate)

}
