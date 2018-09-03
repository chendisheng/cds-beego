package logagent

import(
	"time"
	"fmt"
	"github.com/astaxie/beego/logs"
)


func ServerRun() (err error){
	defer func() {
		err := recover()
		if err != nil {
			logs.Error("error in logagent server run")
			return
		}
	}()

	for {
		//从chan中取出
		msg := GetOneLine()
		fmt.Println(msg)

		err = SendToKafka(msg.Msg, msg.Topic)

		if err != nil {
			fmt.Println("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}
