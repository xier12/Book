package tools

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"strconv"
)

func SendEmail(num int, myemail string) {
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	//         邮件发送方称呼 <发送方的邮箱>
	em.From = "xier12 <1396765025@qq.com>"
	fmt.Println(myemail)
	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{myemail}
	//em.To = []string{"yeqichao265111@qq.com"}
	// 设置主题
	em.Subject = "邮件验证码"
	// 简单设置文件发送的内容，暂时设置成纯文本
	//em.Text = []byte("hello world， 咱们用 golang 发个邮件！！")
	em.Text = []byte(strconv.Itoa(num))
	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1396765025@qq.com", "irnzqssowhzbgagf", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send successfully ... ")

}
