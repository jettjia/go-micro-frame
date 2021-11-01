package gkafka

import "testing"

func Test_ConsumerMessage(t *testing.T) {
	addr := []string{"10.4.7.71:9092"}
	client, err := NewKafkaConsumer(addr)
	if err != nil {
		t.Error(err.Error())
	}

	client.ConsumerMessage("test")
}