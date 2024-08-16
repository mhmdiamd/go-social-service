package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mhmdiamd/go-social-service/internal/config"
	kafkatopic "github.com/mhmdiamd/go-social-service/internal/lib/kafka-topic"
	"github.com/segmentio/kafka-go"
)

type EventPublisherEvent interface {
	PublishCreateEvent(ctx context.Context, ec EventCommittee, topic kafkatopic.EventKafkaEventPublisher) error
}

type KafkaEventPublisherEvent struct {
	writer *kafka.Writer
}

func NewKafkaEventPublisherEvent() *KafkaEventPublisherEvent {
	// host := fmt.Sprintf("%v:%v", config.Cfg.Kafka.Host, config.Cfg.Kafka.Port)
	config := kafka.WriterConfig{
		Brokers: []string{fmt.Sprintf("%v:%v", config.Cfg.Kafka.Host, config.Cfg.Kafka.Port)},
	}

	writer := kafka.NewWriter(config)

	return &KafkaEventPublisherEvent{
		writer: writer,
	}
}

func (kc *KafkaEventPublisherEvent) PublishCreateEvent(ctx context.Context, e EventCommittee, topic kafkatopic.EventKafkaEventPublisher) error {
	// encode data community to json first for sending as topic to kafka
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Value: data,
		Topic: string(topic),
	}

	err = kc.writer.WriteMessages(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
