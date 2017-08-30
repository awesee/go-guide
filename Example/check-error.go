package main

import (
	"fmt"
	. "math"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			panic(err)
		}
	}()
	result := Sqrt(-1)
	fmt.Println("--------------------")
	fmt.Println(result)
	// check(err)
}
