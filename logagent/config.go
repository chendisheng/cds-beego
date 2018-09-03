package logagent

import(
	"fmt"
	"errors"
	"github.com/astaxie/beego"
	"strconv"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath string

	chanSize int
	kafkaAddr string
	collectConf []CollectConf
}

func init(){
	/*定义一个全局变量保存
	var appConfig *Config
	*/
	appConfig = &Config{}
	appConfig.logLevel = beego.AppConfig.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = beego.AppConfig.String("log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs/cds-beego.log"
	}

	size , err := strconv.Atoi(beego.AppConfig.String("log_chan_size"))
	if err != nil {
		appConfig.chanSize = 100
	}else{
		appConfig.chanSize  = size
	}


	appConfig.kafkaAddr = beego.AppConfig.String("log_kafa_address")
	if len(appConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	var cc CollectConf
	cc.LogPath = beego.AppConfig.String("log_collect_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid log_collect_path")
		return
	}

	cc.Topic = beego.AppConfig.String("log_topic")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid log_topic")
		return
	}
	appConfig.collectConf = append(appConfig.collectConf, cc)
	beego.Info("配置项目%v",appConfig)
	/*先测试将log输出配置正确，输出到logagent.log中*/
	beego.Debug("load conf succ, config:%v", appConfig)
}
