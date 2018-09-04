package main

import (
	_ "cds-beego/routers"
	cdslog "cds-beego/logs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init(){
	//日志设置
	initLogConifg()
	//日志代理设置
	//initLogAgent()
}

func initLogConifg(){
	log := logs.NewLogger()
	logs.SetLogger(cdslog.AdapterKafka, `{"address":"192.168.1.22:9092","topic":"cctv1","level":6}`)
	log.Debug("this is a debug message")
}


func main() {

	beego.Run()
}

