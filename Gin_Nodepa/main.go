package main

import (
	"bubble/dao"
	"bubble/log"
	"bubble/models"
	"bubble/routers"
	"bubble/setting"
	"fmt"
)

func main() {

	//if len(os.Args) < 2 {
	//	fmt.Println("Usage：./bubble conf/config.ini")
	//	return
	//}
	// 传入配置文件路径，加载配置文件,
	if err := setting.Init("conf/config.ini"); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	fmt.Println("config.ini配置加载成功", setting.Conf.Port)

	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	fmt.Println("数据库配置初始化加载成功", setting.Conf.MySQLConfig)
	if err := log.InitLogger(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	fmt.Println("日志配置加载成功", setting.Conf.LogConfig)
	log.Logger.Debug("大家好，日志展示")
	defer log.Logger.Sync()
	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
