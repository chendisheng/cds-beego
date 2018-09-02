package logagent

import(
	"github.com/astaxie/beego/logs"
	"time"
	"fmt"
	"cds-beego/logagent/kafka"
	"cds-beego/logagent/tailf"
)

func init(){
	/*
    加载配置文件logagent.conf信息
    */
	filename := "E:\\workspace\\gopath\\src\\cds-beego\\conf\\logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}

	/*
	初始化beego/logs的一些功能，设定输出目录
	*/
	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}

	/*先测试将log输出配置正确，输出到logagent.log中*/
	logs.Debug("load conf succ, config:%v", appConfig)

	/*初始化tailf日志组件 */
	//appConfig.collectConf [{E:\golang\go_pro\logs\logagent.log nginx_log}]
	fmt.Println("appConfig.collectConf",appConfig.collectConf)
	err = tailf.InitTail(appConfig.collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}
	/*先测试将tailf配置正确，输出到logagent.log中*/
	logs.Debug("initialize tailf succ")

	/*初始kafka的工作*/
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	logs.Debug("initialize all succ")

}

func ServerRun() (err error){
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error in logagent server run")
			return
		}
	}()

	for {
		//从chan中取出
		msg := tailf.GetOneLine()
		fmt.Println(msg)

		err = kafka.SendToKafka(msg.Msg, msg.Topic)

		if err != nil {
			fmt.Println("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}
