package main

import "os"

// 下面这个函数演示了如何在过程中使用panic
var user = os.Getenv("USER")

func f() {
	if user == "" {
		panic("no value for $USER")
	}
}

func main() {
	// f()
	throwsPanic(f)
}

// 下面这个函数检查作为其参数的函数在执行时是否会产生panic：

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}
