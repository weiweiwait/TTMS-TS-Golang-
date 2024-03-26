package util

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"time"
)

// 发送邮箱验证码
func SendCode(email string) string {
	//发送对象
	recipient := email
	// 生成验证码
	verificationCode := generateVerificationCode()

	// 构建邮件内容
	subject := "验证码"
	body := fmt.Sprintf("你的验证码是：%s", verificationCode)

	// 创建邮件消息
	message := gomail.NewMessage()
	message.SetHeader("From", "19891294013@163.com")
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	// 创建SMTP客户端
	dialer := gomail.NewDialer("smtp.163.com", 465, "19891294013@163.com", "DRCJMYFWIGGKGSWM")

	// 发送邮件
	err := dialer.DialAndSend(message)
	if err != nil {
		fmt.Println("发送邮件失败:", err)
		return ""
	}

	return verificationCode
}
func generateVerificationCode() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成6位验证码
	code := rand.Intn(899999) + 100000

	// 将验证码转换为字符串
	codeStr := strconv.Itoa(code)

	return codeStr

}
