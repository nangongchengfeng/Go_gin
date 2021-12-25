package main

import (
	"code/mail_qq/log"
	"code/mail_qq/routers"
	"code/mail_qq/setting"
	"fmt"
)

/*
支持多人发送
curl http://10.10.10.3:7070/send -H "Content-Type:application/json" -X POST -d '{"source":"heian","contacts":["账号@mail_qq.com","账号@mail_qq.com"],"subject":"多人测试","content":"现在进行多人测试"}'

*/

/*
zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel
Encoder:编码器(如何写入日志)。我们将使用开箱即用的NewJSONEncoder()
WriterSyncer ：指定日志将写到哪里去。我们使用zapcore.AddSync()
Log Level：哪种级别的日志将被写入。
*/
//var sugarLogger *zap.SugaredLogger

func main() {
	log.InitLogger()
	defer log.SugarLogger.Sync()
	port := setting.GetPort()
	r := routers.SetupRouter()
	r.Run(":" + fmt.Sprint(port))
	log.SugarLogger.Infof("Success! Port is start")
	//r.Run(":8080")

}
