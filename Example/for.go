package main

import "fmt"

func main() {

	// 和 C 语言的 for 一样：
	// for init; condition; post { }
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	// The most basic type, with a single condition.
	// 和 C 的 while 一样：
	i := 10
	for i > 0 {
		fmt.Println(i)
		i--
	}

	// 和 C 的 for(;;) 一样：
	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
