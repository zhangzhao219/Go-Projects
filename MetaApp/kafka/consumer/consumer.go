package consumer

import (
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

type KafkaConsumer struct {
	Node     []string
	Consumer sarama.Consumer
	Topic    string
}

func (c *KafkaConsumer) Consume() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(c.Node, config)
	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}
	defer consumer.Close()

	// 查询该 topic 有多少分区
	partitions, err := consumer.Partitions(c.Topic)
	if err != nil {
		log.Fatal("Partitions err: ", err)
	}
	var wg sync.WaitGroup
	wg.Add(len(partitions))
	// 然后每个分区开一个 goroutine 来消费
	for _, partitionId := range partitions {
		//不开异步会导致一个消费完才会消费另外一个
		go c.consumeByPartition(consumer, c.Topic, partitionId, &wg)
	}
	wg.Wait()
}

func (c *KafkaConsumer) consumeByPartition(consumer sarama.Consumer, topic string, partitionId int32, wg *sync.WaitGroup) {
	defer wg.Done()
	partitionConsumer, err := consumer.ConsumePartition(topic, partitionId, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] Partition = %d, Offset = %d, value = %s\n", message.Partition, message.Offset, string(message.Value))
	}
}
