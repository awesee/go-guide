package main

import (
	"fmt"
	"math"
)

const s string = "constant string"

func main() {
	fmt.Println(s)

	const n = 50000
	const d = 3e10 / n

	fmt.Println(n)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
