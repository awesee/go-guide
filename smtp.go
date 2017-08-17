package main

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
)

func main() {
	host := "smtp.exmail.qq.com"
	email := "mail1@mail.com"
	password := "password"
	toEmail := "mail2@mail.com"
	from := mail.Address{"发送人", email}
	to := mail.Address{"接收人", toEmail}
	body := "我是一封电子邮件!golang发出."

	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte("邮件标题2")))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)
	err := smtp.SendMail(
		host+":25",
		auth,
		email,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		panic(err)
	}
}
