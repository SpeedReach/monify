package infra

import (
	"crypto/tls"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type KafkaWriters struct {
	GroupBill *kafka.Writer
}

func NewKafkaWriter(config Config) (KafkaWriters, error) {
	mechanism, err := scram.Mechanism(scram.SHA256, config.KafkaUsername, config.KafkaPassword)
	if err != nil {
		return KafkaWriters{}, err
	}

	groupBill := &kafka.Writer{
		Topic: "group_bill_modification",
		Addr:  kafka.TCP(config.KafkaConn),
		Transport: &kafka.Transport{
			SASL: mechanism,
			TLS:  &tls.Config{},
		},
	}

	return KafkaWriters{
		GroupBill: groupBill,
	}, nil
}
