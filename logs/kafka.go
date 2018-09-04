package logs

import (
	"log"
	"encoding/json"
	"fmt"
	"time"
	beelog "github.com/astaxie/beego/logs"
	"github.com/Shopify/sarama"
)

// Name for adapter with beego official support
const (
	AdapterKafka   = "kafka"
)
var (
	client sarama.SyncProducer
)


type KAFKAWriter struct{
	Address string `json:"address"`
	Topic   string `json:"topic"`
	Level   int    `json:"level"`
}

// newKAFKAWriter create kafka writer.
func newKAFKAWriter() beelog.Logger {
	return &KAFKAWriter{Level: beelog.LevelTrace}
}

// Init KAFKAWriter with json config string
func (s *KAFKAWriter) Init(jsonconfig string) error {
	err := json.Unmarshal([]byte(jsonconfig), s)
	if err != nil {
		return err
	}
	err = s.initKafka()
	return err
}

// WriteMsg write message in smtp writer.
// it will send message to kafka.
func (s *KAFKAWriter) WriteMsg(when time.Time, msg string, level int) error {
	if level > s.Level {
		return nil
	}
	text := fmt.Sprintf("%s %s", when.Format("2006-01-02 15:04:05"), msg)
	err := s.SendToKafka(text,s.Topic)
	return err
}

/**初始化kafka**/
func(s *KAFKAWriter) initKafka() error{

	addr := s.Address
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	var err error
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		fmt.Errorf("init kafka producer failed, err:", err)
		return err
	}
	return err
}

/*
发送到kafak
*/
func (k *KAFKAWriter)SendToKafka(data, topic string)(err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	_, _, err = client.SendMessage(msg)
	if err != nil{
		fmt .Errorf("send message failed, err:%v data:%v topic:%v", err, data, topic)
		return
	}
	return
}


// Flush implementing method. empty.
func (k *KAFKAWriter) Flush() {
}

// Destroy implementing method. empty.
func (k *KAFKAWriter) Destroy() {
}

func init(){
	log.Println("注册kafka适配器")
	beelog.Register(AdapterKafka, newKAFKAWriter)
}
