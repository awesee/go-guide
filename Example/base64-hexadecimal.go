package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
)

func main() {

	var hex string
	fmt.Print("请输入16进制字符串：")
	fmt.Scanf("%s", &hex)

	buf := new(bytes.Buffer)
	for i, l := 0, len(hex); i < l; i += 2 {
		h := hex[i : i+2]
		num, _ := strconv.ParseUint(h, 16, 0)
		buf.Write([]byte{byte(num)})
	}

	str := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("字符串:", hex)
	fmt.Println("base64编码结果：")
	fmt.Println(str)
}
