package community

import (
	"context"
	"encoding/json"

	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/segmentio/kafka-go"
)

type EventCommunity interface {
	PublishCreateCommunity(ctx context.Context, cm CommunityMember) error
}

type KafkaEventCommunity struct {
	writer *kafka.Writer
}

func NewKafkaEventCommunity() *KafkaEventCommunity {
	config := kafka.WriterConfig{
		Brokers: []string{config.Cfg.Kafka.Host},
	}

	writer := kafka.NewWriter(config)

	return &KafkaEventCommunity{
		writer: writer,
	}
}

func (kc *KafkaEventCommunity) PublishCreateCommunity(ctx context.Context, cm CommunityMember, topic string) error {
	// encode data community to json first for sending as topic to kafka
	data, err := json.Marshal(cm)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Value: data,
		Topic: topic,
	}

	err = kc.writer.WriteMessages(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
