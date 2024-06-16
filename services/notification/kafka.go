package notification

import (
	"crypto/tls"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"time"
)

type KafkaConsumers struct {
	GroupBillsConsumer *kafka.Reader
}

func NewConsumers(config Config) (KafkaConsumers, error) {
	mechanism, err := scram.Mechanism(scram.SHA256, config.KafkaUser, config.KafkaPassword)
	if err != nil {
		return KafkaConsumers{}, err
	}

	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.KafkaConn},
		Topic:   "group_bill_modification",
		Dialer:  dialer,
		GroupID: "notification-group",
	})

	return KafkaConsumers{
		GroupBillsConsumer: r,
	}, nil
}
