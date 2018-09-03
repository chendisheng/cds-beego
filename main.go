package main

import (
	_ "cds-beego/routers"
	_ "github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"cds-beego/logagent"

)

func init(){
	//日志设置
	initLogConifg()
	//日志代理设置
	initLogAgent()
}

func initLogConifg(){
	log := logs.NewLogger()
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/cds-beego.log"}`)
	log.Debug("this is a debug message")
}

func initLogAgent(){
	beego.Info("运行日志代理....")
	go logagent.ServerRun()

}

func main() {

	beego.Run()
}

