package communitymember

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mhmdiamd/go-social-service/domain/community"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/mhmdiamd/go-social-service/internal/log"
	"github.com/segmentio/kafka-go"
)

type EventReaderCommunityMember interface {
	ReadCreateCommunity(ctx context.Context, topic string, handler func(ctx context.Context, cm AddCommunityMemberRequestPayload) error) error
	DeleteCommunityById(ctx context.Context, topic string, handler func(ctx context.Context, community_id int) error) error
	Init(ctx context.Context) error
}

type KafkaEventReaderCommunityMember struct {
	groupID string
	svc     service
	offsets map[int]int64 // Map to store committed offsets (replace with appropriate data structure)
}

func NewKafkaReaderConfig(groupID string, topic string) kafka.ReaderConfig {
	// host := fmt.Sprintf("%s:%s", config.Cfg.Kafka.Host, config.Cfg.Kafka.Port)
	return kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", config.Cfg.Kafka.Host, config.Cfg.Kafka.Port)},
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

func (kc *KafkaEventReaderCommunityMember) Init(ctx context.Context) error {
	reader := NewEventReaderCommunityMember(kc.groupID, kc.svc)

	if err := reader.ReadCreateCommunity(context.Background(), community.CREATE_COMMUNITY, kc.svc.AddMember); err != nil {
		return err
	}

	if err := reader.DeleteCommunityById(context.Background(), community.DELETE_COMMUNITY_BY_ID, kc.svc.DeleteCommunityMemberByCommunityId); err != nil {
		return err
	}

	return nil
}

func (kc *KafkaEventReaderCommunityMember) ReadCreateCommunity(ctx context.Context, topic string, handler func(ctx context.Context, cm AddCommunityMemberRequestPayload) error) error {
	// encode data community to json first for sending as topic to kafka
	config := NewKafkaReaderConfig(kc.groupID, topic)
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

func (kc *KafkaEventReaderCommunityMember) DeleteCommunityById(ctx context.Context, topic string, handler func(ctx context.Context, communityId int) error) error {
	config := NewKafkaReaderConfig(kc.groupID, topic)
	reader := kafka.NewReader(config)

	defer reader.Close()

	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		log.Log.Errorf(ctx, "[DeleteCommunityById (event reader), reader.ReadMessage (kafka)] with error detail %s", err.Error())
		return err
	}

	var data struct{ id int }
	if err := json.Unmarshal(msg.Value, &data); err != nil {
		log.Log.Errorf(ctx, "[DeleteCommunityById (event reader), reader.ReadMessage (kafka)] with error detail %s", err.Error())
		return err
	}

	if err := handler(ctx, data.id); err != nil {
		log.Log.Errorf(ctx, "[DeleteCommunityById (event reader), handler (service)] with error detail %s", err.Error())
		return err
	}

	return err
}
