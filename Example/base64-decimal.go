package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

func main() {

	var num rune
	fmt.Print("请输入数字：")
	fmt.Scanf("%d", &num)
	// fmt.Println("数字: ", num)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, num)
	// fmt.Println("bytes: ", buf.Bytes())

	bt := bytes.TrimLeftFunc(buf.Bytes(), func(r rune) bool {
		return r == rune(0)
	})
	// fmt.Println("bt: ", bt)

	str := base64.StdEncoding.EncodeToString(bt)
	fmt.Println("数字:", num, ", base64编码结果：")
	fmt.Println(str)
}
