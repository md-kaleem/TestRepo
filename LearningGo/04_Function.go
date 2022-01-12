package main

import "fmt"

func calc(num1 int, num2 int) (out1, out2 int) {
	out1 = num1 + num2
	out2 = num1 - num2
	return
}

func main() {
	result1, result2 := calc(4, 8)
	fmt.Println(result1, result2)
}
