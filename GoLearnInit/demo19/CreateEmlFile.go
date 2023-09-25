package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	to := "recipient@example.com" // 收件人邮箱地址
	cc := "cc@example.com"        // 抄送人邮箱地址
	subject := "Test Email"       // 邮件主题
	htmlContent := `<html>
	<head></head>
	<body>
	<h1>Hello, this is a test email</h1>
	<p>This email contains <strong>formatted</strong> text.</p>
	</body>
	</html>` // 邮件内容（HTML格式）

	// 生成eml文件内容
	emlContent := createEMLContent(to, cc, subject, htmlContent)

	// 将eml内容写入文件
	err := ioutil.WriteFile("CreateEmlFile.eml", []byte(emlContent), 0644)
	if err != nil {
		fmt.Println("生成eml文件失败:", err)
	} else {
		fmt.Println("eml文件生成成功")
	}
}

// 生成eml文件内容
func createEMLContent(to, cc, subject, htmlContent string) string {
	// 构建eml头部
	header := make(map[string]string)
	header["To"] = to
	header["Cc"] = cc
	header["Subject"] = subject
	header["Content-Type"] = "text/html; charset=UTF-8"

	// 构建eml内容
	emlContent := ""
	for key, value := range header {
		emlContent += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	emlContent += "\r\n" + htmlContent

	return emlContent
}
