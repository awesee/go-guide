package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println("empty:", arr)

	arr[5] = 55
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[5])
	// The builtin len returns the length of an array.
	fmt.Println("len:", len(arr))

	// Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
