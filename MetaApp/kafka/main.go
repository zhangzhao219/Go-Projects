package main

import (
	"fmt"
	"kafkaDemo/consumer"
	"kafkaDemo/producer"
	"log"
	"sync"
	"time"
)

func main() {
	Conn := "127.0.0.1:9092"
	topic := "test_log"

	var wg sync.WaitGroup
	wg.Add(2)

	// 消费者
	go func() {
		defer wg.Done()
		// 初始化consumer
		var kafkaConsumer = consumer.KafkaConsumer{
			Node:  []string{Conn},
			Topic: topic,
		}
		// 消费
		go kafkaConsumer.Consume()
	}()

	// 生产者
	go func() {
		defer wg.Done()

		index := 0
		for {
			// 生产者发送消息
			_, err := producer.Send(Conn, topic, fmt.Sprintf("lox_%d", index))
			if err != nil {
				log.Print("测试失败:" + err.Error())
				return
			}
			index++
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}
