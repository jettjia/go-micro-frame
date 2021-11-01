package gkafka

import (
	"fmt"
	"testing"
)

// 发送消息
func TestKafkaProducer_ProducerMessage(t *testing.T) {
	addr := []string{"10.4.7.71:9092"}
	client, err := NewKafkaProducer(addr)
	fmt.Println(client)
	fmt.Println("-------- err ",err)
	if err != nil {
		t.Error(err.Error())
	}

	_ = client.ProducerMessage("test", "www.baidu.com")
}