package communitymember

import (
	"context"
	"encoding/json"

	"github.com/mhmdiamd/go-social-service/domain/community"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/segmentio/kafka-go"
)

type EventReaderCommunityMember interface {
	ReadCreateCommunity(ctx context.Context, topic string, handler func(ctx context.Context, cm AddCommunityMemberRequestPayload) error) error
}

type KafkaEventReaderCommunityMember struct {
	groupID string
	svc     service
	offsets map[int]int64 // Map to store committed offsets (replace with appropriate data structure)
}

func NewKafkaConfig(groupID string, topic string) kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers: []string{config.Cfg.Kafka.Host},
		GroupID: groupID,
		Topic:   topic,
	}
}

func NewEventReaderCommunityMember(groupID string, svc service) EventReaderCommunityMember {
	return &KafkaEventReaderCommunityMember{
		svc:     svc,
		groupID: groupID,
	}
}

func (kc *KafkaEventReaderCommunityMember) ReadCreateCommunity(ctx context.Context, topic string, handler func(ctx context.Context, cm AddCommunityMemberRequestPayload) error) error {
	// encode data community to json first for sending as topic to kafka
	config := NewKafkaConfig(kc.groupID, topic)
	reader := kafka.NewReader(config)

	defer reader.Close()

	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		return err
	}

	var cm community.CommunityMember
	err = json.Unmarshal(msg.Value, &cm)
	if err != nil {
		return err
	}

	req := AddCommunityMemberRequestPayload{
		Role:         CommunityMemberRole(cm.Role),
		CommunityId:  cm.CommunityId,
		UserPublicId: cm.UserPublicId.String(),
	}

	if err := handler(ctx, req); err != nil {
		return err
	}

	return nil
}
