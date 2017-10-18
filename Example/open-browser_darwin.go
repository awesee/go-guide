package main

import "os/exec"

func main() {
	exec.Command("open", "http://www.baidu.com").Start()
}
