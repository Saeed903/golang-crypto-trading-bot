package kafka

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

var (
	brokers = []string{"127.0.0.1:9092"}
	topic   = "BTC-USDT"
)

func newKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0
	return conf
}

func newKafkaSyncProducer() sarama.SyncProducer {
	syncProducer, err := sarama.NewSyncProducer(brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}
	return syncProducer
}

func newKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("kafka error: %s\n", err)
		os.Exit(-1)
	}
	return consumer
}

func sendMessage(syncProducer sarama.SyncProducer, event interface{}) error {
	json, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msglog := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(json)),
	}

	partition, offset, err := syncProducer.SendMessage(msglog)
	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
	}

	fmt.Printf("Message: %+v\n", event)
	fmt.Printf("Message is stored in partition %d, offset %d\n", partition, offset)

	return nil
}
