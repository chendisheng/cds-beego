package logagent

import (
	"github.com/hpcloud/tail"
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"strconv"
)

type CollectConf struct {
	LogPath string
	Topic   string
}
/*{E:\golang\go_pro\logs\logagent.log nginx_log}
每条配置
*/
type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}

type TextMsg struct {
	Msg string
	Topic string
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan chan *TextMsg
}

var (
	tailObjMgr* TailObjMgr
)

func init(){
	 conf := appConfig.collectConf
	if len(conf) == 0 {
		logs.Error("invalid config for log collect, conf:%v", conf)
		return
	}

	chanSizeStr := beego.AppConfig.String("log_chan_size")
	chanSize ,err := strconv.Atoi(chanSizeStr)
    if err != nil {
    	return
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan*TextMsg, chanSize),
	}
	////appConfig.collectConf [{E:\golang\go_pro\logs\logagent.log nginx_log}]
	for _, v := range conf {
		obj := &TailObj{
			conf: v,
		}
		//v--- {E:\golang\go_pro\logs\logagent.log nginx_log}
		fmt.Println("v---",v)
		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})

		if errTail != nil {
			err = errTail
			return
		}

		obj.tail = tails
		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

		go readFromTail(obj)
	}
}

func GetOneLine()(msg *TextMsg) {
	msg = <- tailObjMgr.msgChan
	return
}

func readFromTail(tailObj *TailObj) {
	for true {
		line, ok := <-tailObj.tail.Lines
		if !ok {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		textMsg := &TextMsg{
			Msg:line.Text,
			Topic: tailObj.conf.Topic,
		}
		tailObjMgr.msgChan <- textMsg
	}
}
