package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	d := http.Dir("/")
	fmt.Println(d)
	http.Handle("/", http.FileServer(http.Dir("/")))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
