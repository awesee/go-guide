package main

import "os/exec"

func main() {
	exec.Command("start", "iexplore", "http://www.baidu.com").Start()
}
