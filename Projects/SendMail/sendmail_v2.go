package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
支持多人发送
curl http://10.10.10.3:7070/send -H "Content-Type:application/json" -X POST -d '{"source":"heian","contacts":["账号@qq.com","账号@qq.com"],"subject":"多人测试","content":"现在进行多人测试"}'

*/
// 定义接收数据的结构体
type User struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Source   string   `form:"source" json:"source" uri:"source" xml:"source" binding:"required"`
	Contacts []string `form:"contacts" json:"contacts" uri:"contacts" xml:"contacts" binding:"required"`
	Subject  string   `form:"subject" json:"subject" uri:"subject" xml:"subject" binding:"required"`
	Content  string   `form:"content" json:"content" uri:"content" xml:"content" binding:"required"`
}

func SendToMail(user, sendUserName, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	//fmt.Println(hp)
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	//fmt.Println(err)
	return err
}

func PostMail(c *gin.Context) {
	//// 声明接收的变量
	var json User
	//// 将request的body中的数据，自动按照json格式解析到结构体
	//
	if err := c.ShouldBindJSON(&json); err != nil {
		//	// 返回错误信息
		//	// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println(json.Content, json.Contacts)
	c.JSON(http.StatusOK, gin.H{"status": &json})
	user := "账号@qq.com"
	password := "密码"
	host := "smtp.qq.com:25"
	source := json.Source
	if source != "heian" {
		fmt.Println("Send mail error!,source 认证失败")
		c.JSON(http.StatusOK, gin.H{
			"error": "Send mail error!,source 认证失败",
		})
		return
	}
	//println(json.Contacts)
	to := json.Contacts
	//if strings.TrimSpace(to) == "" {
	//	fmt.Println("Send mail error!,发送人为空")
	//	c.JSON(http.StatusOK, gin.H{
	//		"error": "Send mail error!,发送人为空",
	//	})
	//	return
	//}
	subject := json.Subject
	if strings.TrimSpace(subject) == "" {
		fmt.Println("Send mail error!标题为空")
		c.JSON(http.StatusOK, gin.H{
			"error": "Send mail error!,标题为空",
		})
		return
	}
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>
			` + fmt.Sprintf(json.Content) +
		`</body>
		</html>`

	sendUserName := "告警平台" //发送邮件的人名称
	fmt.Println("send email")

	for _, s := range to {
		//fmt.Println(i, s)
		err := SendToMail(user, sendUserName, password, host, s, subject, body, "html")
		//log.Printf("接收人：", s+"\n"+"标题:", json.Subject+"\n", "发送内容：", json.Content+"\n")
		fmt.Printf("接收人:%s \n 标题: %s \n 内容: %s \n", s, json.Subject, json.Content)
		if err != nil {
			fmt.Println("Send mail error!")
			c.JSON(http.StatusOK, gin.H{
				"error": "Send mail error! !",
			})
			//fmt.Println(err)
		} else {
			fmt.Println("Send mail success!")
			c.JSON(http.StatusOK, gin.H{
				"success": "Send mail success! !",
			})
		}

	}
	//err := SendToMail(user, sendUserName, password, host, string(to), subject, body, "html")
	//if err != nil {
	//	fmt.Println("Send mail error!")
	//	c.JSON(http.StatusOK, gin.H{
	//		"error": "Send mail error! !",
	//	})
	//	//fmt.Println(err)
	//} else {
	//	fmt.Println("Send mail success!")
	//	c.JSON(http.StatusOK, gin.H{
	//		"success": "Send mail success! !",
	//	})
	//}

}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// JSON绑定
	r.POST("send", PostMail)
	r.Run(":7070")
}
