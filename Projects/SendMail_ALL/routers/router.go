package routers

import (
	"code/mail_qq/app"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//启动连接数据库，根据用户请求
	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/send", app.PostMail)

	}

	return r
}
