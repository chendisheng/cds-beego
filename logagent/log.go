package logagent

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

/*
初始化beego/logs的一些功能,设定输出目录
*/
func init(){
	config := make(map[string]interface{})
	config["filename"] = appConfig.logPath
	config["level"] = convertLogLevel(appConfig.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		logs.Error("initLogger failed, marshal err:",err)
		return
	}
	//logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.Info("initLogger config:",string(configStr))

}

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
