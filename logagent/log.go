package logagent

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {

	switch (level) {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return  logs.LevelDebug
}

/*
    初始化beego/logs的一些功能,设定输出目录
*/
func initLogger()(err error) {

	config := make(map[string]interface{})
	config["filename"] = appConfig.logPath
	config["level"] = convertLogLevel(appConfig.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	//{"filename":"E:\\golang\\go_pro\\logs\\logagent.log","level":7}
	fmt.Println(string(configStr))
	return
}
