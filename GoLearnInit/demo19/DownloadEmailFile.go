package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 定义邮件的主题、收件人、抄送人和内容
	subject := "Test email"
	to := "receiver@example.com"
	cc := "cc@example.com"
	body := `
        <html>
            <head></head>
            <body>
                <h1>Hello, this is a test email</h1>
                <p>This email contains <strong>formatted</strong> text.</p>
            </body>
        </html>
    `

	// 将主题、收件人、抄送人和内容组合成一个字符串
	emailContent := fmt.Sprintf(`Subject: %s
To: %s
Cc: %s

	%s`, subject, to, cc, body)

	// 定义处理下载请求的路由函数
	http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头，指定文件名和 MIME 类型
		w.Header().Set("Content-Disposition", "attachment; filename=email.eml")
		w.Header().Set("Content-Type", "application/octet-stream")

		// 将邮件内容写入响应体
		_, err := w.Write([]byte(emailContent))
		if err != nil {
			log.Fatal(err)
		}
	})

	// 启动 HTTP 服务器
	log.Fatal(http.ListenAndServe(":8080", nil))
}
