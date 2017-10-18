package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	port := 9090
	ip := getIp()
	cmd := exec.Command("open", "http://"+ip+":"+strconv.Itoa(port))
	cmd.Start()

	err := http.ListenAndServe(":"+strconv.Itoa(port), http.FileServer(http.Dir(".")))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func getIp() (ip string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}

		}
	}

	return
}
