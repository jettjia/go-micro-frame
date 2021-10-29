package rabbitmq

import (
	"fmt"
)

type Registry struct {
}

type RegistryClient interface {
	Register(host string, port int, user string, password string) string
	DoPublish(amqpURI, exchange, exchangeType, routingKey, body string, reliable bool) error
	DoConsumer(amqpURI, exchange, exchangeType, routingKey, queue string) (*Consumer, error)
}

func NewRegistryClient() RegistryClient {
	return &Registry{
	}
}

// 获取rabbitMq 连接地址
func (r *Registry) Register(host string, port int, user string, password string) string {
	//"amqp://admin:123456@10.4.7.71:5672/"
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port)
	return uri
}

// 发送消息
func (r *Registry) DoPublish(amqpURI, exchange, exchangeType, routingKey, body string, reliable bool) error {
	return Publish(amqpURI, exchange, exchangeType, routingKey, body, reliable)
}

// 消费消息
func (r *Registry) DoConsumer(amqpURI, exchange, exchangeType, routingKey, queueName string) (*Consumer, error) {
	consumer, err := NewConsumer(amqpURI, exchange, exchangeType, queueName, routingKey, "")
	return consumer, err
}
