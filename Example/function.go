package main

import "fmt"

func plus(num1, num2 int) int {
	return num1 + num2
}

func main() {
	sum := plus(1, 2)
	fmt.Println("1 + 2 =", sum)
}
