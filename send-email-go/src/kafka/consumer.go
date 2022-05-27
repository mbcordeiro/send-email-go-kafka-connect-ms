package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumer struct {
	Topics    []string
	ConfigMap *ckafka.ConfigMap
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		Topics:    topics,
		ConfigMap: configMap,
	}
}
