package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Release      bool `ini:"release"`
	Port         int  `ini:"port"`
	*MySQLConfig `ini:"mysql"`
	*LogConfig   `ini:"log"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

type LogConfig struct {
	Level      string `ini:"level"`
	Filename   string `ini:"filename"`
	MaxSize    int    `ini:"maxsize"`
	MaxAge     int    `ini:"max_age"`
	MaxBackups int    `ini:"max_backups"`
}

//把先关初始的参数加载到全局变量，然后方便调用
func Init(file string) error {
	return ini.MapTo(Conf, file)
}
