package community

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/segmentio/kafka-go"
)

const (
	CREATE_COMMUNITY       = "CREATE_COMMUNITY"
	DELETE_COMMUNITY_BY_ID = "DELETE_COMMUNITY_BY_ID"
)

type KafkaEventPublisherCommunity struct {
	writer *kafka.Writer
}

func NewKafkaEventPublisherCommunity() *KafkaEventPublisherCommunity {
	// host := fmt.Sprintf("%v:%v", config.Cfg.Kafka.Host, config.Cfg.Kafka.Port)
	config := kafka.WriterConfig{
		Brokers: []string{fmt.Sprintf("%v:%v", config.Cfg.Kafka.Host, config.Cfg.Kafka.Port)},
	}

	writer := kafka.NewWriter(config)

	return &KafkaEventPublisherCommunity{
		writer: writer,
	}
}

func (kc *KafkaEventPublisherCommunity) PublishCreateCommunity(ctx context.Context, cm CommunityMember, topic string) error {
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

func (kc *KafkaEventPublisherCommunity) PublishDeleteCommunityById(ctx context.Context, communityId int, topic string) error {
	// encode data community to json first for sending as topic to kafka
	data, err := json.Marshal(map[string]int{
		"community_id": communityId,
	})
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
