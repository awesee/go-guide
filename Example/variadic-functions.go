package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2, 3)
	sum(2, 4, 6)

	nums := []int{1, 3, 5}
	sum(nums...)
}
