package gkafka

import (
	"fmt"
	
	"github.com/Shopify/sarama"
)

type KafkaConsumer struct {
	Addr          []string
	KafkaConsumer sarama.Consumer
}

func NewKafkaConsumer(addr []string) (*KafkaConsumer, error) {
	c := &KafkaConsumer{}

	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		return nil, err
	}
	c.Addr = addr
	c.KafkaConsumer = consumer

	return c, nil
}

func (c *KafkaConsumer) ConsumerMessage(topic string) {
	partitionList, err := c.KafkaConsumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		return
	}

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := c.KafkaConsumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
}
