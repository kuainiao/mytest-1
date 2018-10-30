package 发送邮件

import (
	"testing"
		"net/smtp"
	"fmt"
	"strings"
)

func TestA(t *testing.T) {
	a()
}

//
//func a() {
//	auth := smtp.PlainAuth("", "zhaojinhui9988@163.com", "jszx_zjj2016", "smtp.163.com")
//
//	to := []string{"zhaoji@ifeng.com"}
//	//image,_:=ioutil.ReadFile("/home/hero/Pictures/2018-10-24 16-28-26 的屏幕截图.png")
//	//imageBase64:=base64.StdEncoding.EncodeToString(image)
//	msg := []byte("from:xxx@163.com\r\n"+
//		"to: xxx@163.com\r\n" +
//		"Subject: hello,subject!\r\n"+
//		"Content-Type:multipart/mixed;boundary=a\r\n"+
//		"Mime-Version:1.0\r\n"+
//		"\r\n" +
//		"--a\r\n"+
//		"Content-type:text/plain;charset=utf-8\r\n"+
//		"Content-Transfer-Encoding:quoted-printable\r\n"+
//		"\r\n"+
//		"此处为正文内容!\r\n"+
//		"--a\r\n"+
//		"Content-type:image/jpg;name=1.jpg\r\n"+
//		"Content-Transfer-Encoding:base64\r\n"+
//		"\r\n"+
//		//imageBase64+"\r\n"+
//		"--a--\r\n")
//	err := smtp.SendMail("smtp.163.com:25", auth, "zhaojinhui9988@163.com", to, msg)
//	if err != nil {
//		fmt.Println(err)
//	}
//}







func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func a() {
	user := "zhaoji@ifeng.com"
	password := "ifeng_zjj2018"
	host := "mail.ifeng.com:25"
	to := "784314557@qq.com"

	subject := "Test send email by golang"

	body := `
<html>
<body>
<h3>
"Test send email by golang"
</h3>
</body>
</html>
`
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
}
