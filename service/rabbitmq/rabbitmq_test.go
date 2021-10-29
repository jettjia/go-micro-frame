package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"testing"
	"time"
)

// 生产者
func Test_DoPublish(t *testing.T) {
	registerClient := NewRegistryClient()
	amqpURI := registerClient.Register("10.4.7.71", 5672, "admin", "123456")

	// 发送消息测试
	err := registerClient.DoPublish(amqpURI, "test-exchange-micro", "direct", "test-routingKey-micro", "id=100", true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("rabbitMq publish success")
}

//消费者
func Test_Consumer(t *testing.T) {
	registerClient := NewRegistryClient()
	amqpURI := registerClient.Register("10.4.7.71", 5672, "admin", "123456")
	exchange := "test-exchange-micro"
	exchangeType := "direct"
	routingKey := "test-routingKey-micro"
	queueName := "test-queue-micro"

	// 处理消费逻辑
	{
		c := &Consumer{
			conn:    nil,
			channel: nil,
			done:    make(chan error),
		}

		var err error

		c.conn, err = amqp.Dial(amqpURI)
		if err != nil {
			t.Error(err)
		}

		go func() {
			fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
		}()

		fmt.Printf("got Connection, getting Channel")
		c.channel, err = c.conn.Channel()
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("got Channel, declaring Exchange (%q)", exchange)
		if err = c.channel.ExchangeDeclare(
			exchange,     // name of the exchange
			exchangeType, // type
			true,         // durable
			false,        // delete when complete
			false,        // internal
			false,        // noWait
			nil,          // arguments
		); err != nil {
			t.Error(err)
		}

		fmt.Printf("declared Exchange, declaring Queue %q", queueName)
		queue, err := c.channel.QueueDeclare(
			queueName, // name of the queue
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // noWait
			nil,       // arguments
		)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
			queue.Name, queue.Messages, queue.Consumers, routingKey)

		if err = c.channel.QueueBind(
			queue.Name, // name of the queue
			routingKey, // bindingKey
			exchange,   // sourceExchange
			false,      // noWait
			nil,        // arguments
		); err != nil {
			t.Error(err)
		}

		fmt.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
		deliveries, err := c.channel.Consume(
			queue.Name, // name
			c.tag,      // consumerTag,
			false,      // noAck
			false,      // exclusive
			false,      // noLocal
			false,      // noWait
			nil,        // arguments
		)
		if err != nil {
			t.Error(err)
		}

		go doHandle(deliveries, c.done)
	}
}

func doHandle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		log.Printf(
			"获取到的具体内容是： %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}

// 在consumer直接处理了，不适用此方式
func Test_DoConsumer(t *testing.T) {
	registerClient := NewRegistryClient()
	amqpURI := registerClient.Register("10.4.7.71", 5672, "admin", "123456")
	exchange := "test-exchange-micro"
	exchangeType := "direct"
	routingKey := "test-routingKey-micro"
	queueName := "test-queue-micro"

	{
		// 获取消息
		c, err := registerClient.DoConsumer(amqpURI, exchange, exchangeType, routingKey, queueName)
		if err != nil {
			t.Fatal(err)
		}

		// 时间
		lifetime := time.Duration(0) * time.Second
		if lifetime > 0 {
			log.Printf("running for %s", lifetime)
			time.Sleep(lifetime)
		} else {
			log.Printf("running forever")
			select {}
		}
		log.Printf("shutting down")

		if err := c.Shutdown(); err != nil {
			log.Fatalf("error during shutdown: %s", err)
		}
	}

	fmt.Println("rabbitMq consumer success")
}
