package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 定义邮件的主题、收件人和内容
	// 定义邮件的主题、收件人、抄送人和内容
	subject := "Test email"
	to := "receiver@example.com"
	cc := "cc@example.com"
	body := `<html>
	<head></head>
	<body>
	<h1>Hello, this is a test email</h1>
	<p>This email contains <strong>formatted</strong> text.</p>
	</body>
	</html>` // 邮件内容（HTML格式）

	// 将主题、收件人、抄送人和内容组合成一个字符串
	emailContent := fmt.Sprintf(`Subject: %s
To: %s
Cc: %s
Content-Type: text/html; charset=UTF-8

    %s`, subject, to, cc, body)

	// 打开一个名为 email.eml 的文件，如果不存在则创建它
	file, err := os.OpenFile("CreateEmailFile.eml", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 将邮件内容写入文件
	err1 := ioutil.WriteFile(file.Name(), []byte(emailContent), 0644)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("Email file created successfully.")
}
