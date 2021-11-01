package gkafka

import (
	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	Addr     []string
	producer sarama.SyncProducer
}

func NewKafkaProducer(addr []string) (*KafkaProducer, error) {
	sp := &KafkaProducer{}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true

	var err error
	if sp.producer, err = sarama.NewSyncProducer(addr, config); err != nil {
		return nil, err
	}

	sp.Addr = addr

	return sp, nil
}

// 发送消息
func (g *KafkaProducer) ProducerMessage(topic string, data string) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}

	_, _, err := g.producer.SendMessage(message)

	return err
}

func (g *KafkaProducer) Close() error {
	return g.producer.Close()
}
