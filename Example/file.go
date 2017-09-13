package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	if arg == "" {
		fmt.Println("请输入服务地址")
		os.Exit(1)
	}
	port := flag.Arg(1)
	if port == "" {
		port = "88"
	}
	http.Handle("/", http.FileServer(http.Dir(arg)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
