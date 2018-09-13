package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	var num int64
	fmt.Print("请输入数字：")
	fmt.Scanf("%d", &num)

	// Method 1
	// buf := new(bytes.Buffer)
	// binary.Write(buf, binary.BigEndian, num)
	// bt := bytes.TrimLeftFunc(buf.Bytes(), func(r rune) bool {
	// 	return r == rune(0)
	// })

	// Method 2
	fmt.Println("数字:", num, ", base64编码结果：")
	bt := []byte{byte(num % 256)}
	for num >>= 8; num > 0; num >>= 8 {
		bt = append([]byte{byte(num % 256)}, bt...)
	}
	str := base64.StdEncoding.EncodeToString(bt)
	fmt.Println(str)
}
