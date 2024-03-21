package producer

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

// 随机发送
func Send(conn, topic, content string) (bool, error) {

	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	SyncProducer, err := sarama.NewSyncProducer(
		[]string{conn},
		config,
	)

	if err != nil {
		return false, nil
	}

	// 构建发送的消息
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(time.Now().String()),
		Value: sarama.StringEncoder(content),
	}

	// SendMessage：该方法是生产者生产给定的消息
	// 生产成功的时候返回该消息的分区和所在的偏移量
	// 生产失败的时候返回error
	partition, offset, err := SyncProducer.SendMessage(msg)
	if err != nil {
		return false, err
	}
	log.Printf("[Producer] Partition = %d, Offset = %d, value = %s\n", partition, offset, content)
	return true, nil
}
