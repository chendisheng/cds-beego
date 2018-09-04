package logagent

import(
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"fmt"
)

var (
	client sarama.SyncProducer
)

/**初始化kafka**/
func init(){
	addr :=beego.AppConfig.String("log_kafa_address")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
    var err error
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka producer failed, err:", err)
		return
	}
	//记录步骤信息
	logs.Debug("init kafka succ")
	return
}

/*
    发送到kafak
*/
func SendToKafka(data, topic string)(err error) {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err:%v data:%v topic:%v", err, data, topic)
		return
	}

	fmt.Println("send succ, pid:%v offset:%v, topic:%v\n", pid, offset, topic)
	return
}
