package logagent

import(
	"fmt"
	"errors"
	"github.com/astaxie/beego/config"
	"cds-beego/logagent/tailf"
	"github.com/astaxie/beego"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath string

	chanSize int
	kafkaAddr string
	collectConf []tailf.CollectConf
}

func loadCollectConf(conf config.Configer) (err error ) {

	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}

	appConfig.collectConf = append(appConfig.collectConf, cc)
	beego.Info("配置项目%v",appConfig)
	return
}

/*
    加载配置文件信息
    [logs]
    log_level=debug
    log_path=E:\golang\go_pro\logs\logagent.log
    [collect]
    log_path=E:\golang\go_pro\logs\logagent.log
    topic=nginx_log

    chan_size=100
    [kafka]
    server_addr=192.168.21.8:9092
*/
func loadConf(confType, filename string) (err error) {

	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	/*定义一个全局变量保存
	var appConfig *Config
	*/
	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "E:\\workspace\\gopath\\src\\cds-beego\\logs\\logagent.log"
	}

	appConfig.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}

	appConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed, err:%v\n", err)
		return
	}
	return
}