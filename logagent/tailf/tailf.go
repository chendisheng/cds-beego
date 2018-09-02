package tailf

import (
	"github.com/hpcloud/tail"
	_ "github.com/astaxie/beego/logs"
	"fmt"
	"time"
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

func GetOneLine()(msg *TextMsg) {
	msg = <- tailObjMgr.msgChan
	return
}
/*初始化Tail组件一些功能*/
func InitTail(conf []CollectConf, chanSize int) (err error) {

	if len(conf) == 0 {
		err = fmt.Errorf("invalid config for log collect, conf:%v", conf)
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

	return
}

func readFromTail(tailObj *TailObj) {
	for true {
		line, ok := <-tailObj.tail.Lines
		if !ok {
			//logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
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
