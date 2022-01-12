package main

import "fmt"

var a = 9 // package level variable

func demo() {
	var a = 8 // function level variable
	fmt.Println("in demo", a)
}

func main() {
	demo()
	fmt.Println("in main", a)
}
