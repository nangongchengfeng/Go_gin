package setting

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

//import (
//	"github.com/go-ini/ini"
//)
//
//var Conf = new(AppConfig)
//
//// AppConfig 应用程序配置
//type AppConfig struct {
//	Release     bool `ini:"release"`
//	Port        int  `ini:"port"`
//	*MailConfig `ini:"mail"`
//}
//
//// MailConfig 数据库配置
//type MailConfig struct {
//	User     string `ini:"user"`
//	Password string `ini:"password"`
//	Source   string `ini:"source"`
//	Host     string `ini:"host"`
//	Port     int    `ini:"port"`
//}
//
//func Init(file string) error {
//	return ini.MapTo(Conf, file)
//}

func GetMail() (user, password, host, source string) {
	//读取.ini里面的数据库配置
	config, err := ini.Load("conf/config.ini")
	if err != nil {
		//失败
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	host = config.Section("mail").Key("host").String()
	//port = config.Section("mail").Key("port").String()
	user = config.Section("mail").Key("user").String()
	password = config.Section("mail").Key("password").String()
	source = config.Section("mail").Key("source").String()
	//fmt.Println(user, password, host, source)
	return
}
func GetPort() (prot string) {
	config, err := ini.Load("conf/config.ini")
	if err != nil {
		//失败
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	prot = config.Section("").Key("port").String()
	return
}
