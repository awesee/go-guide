package main

import "fmt"

func main() {

	// We’ll iterate over 2 values in the queue channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	fmt.Println("======")

	// This range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
