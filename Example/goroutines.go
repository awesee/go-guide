package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}

}

func main() {

	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行

	// f("direct")

	// go f("goroutine")

	// go func(msg string) {
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println(msg, ":", i)
	// 	}
	// }("going")

	// var input string
	// fmt.Scanln(&input)
	// fmt.Println("done")
}
