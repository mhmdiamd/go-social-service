package eventcommittee

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mhmdiamd/go-social-service/internal/config"
	kafkatopic "github.com/mhmdiamd/go-social-service/internal/lib/kafka-topic"
	"github.com/mhmdiamd/go-social-service/internal/log"
	"github.com/segmentio/kafka-go"
)

type EventReaderEventCommittee interface {
	ReadCreateEventCommittee(ctx context.Context, topic kafkatopic.EventKafkaEventPublisher, handler func(ctx context.Context, payload CreateEventCommitteeRequestPayload) error) error
	Init(ctx context.Context) error
}

type KafkaEventReaderEventCommittee struct {
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

func NewEventReaderEventCommittee(groupID string, svc service) *KafkaEventReaderEventCommittee {
	return &KafkaEventReaderEventCommittee{
		svc:     svc,
		groupID: groupID,
	}
}

func (kc *KafkaEventReaderEventCommittee) Init(ctx context.Context) error {
	reader := NewEventReaderEventCommittee(kc.groupID, kc.svc)

	if err := reader.ReadCreateEventCommittee(context.Background(), kafkatopic.CREATE_EVENT_COMMITTEE, kc.svc.AddEventCommittee); err != nil {
		return err
	}

	return nil
}

func (kc *KafkaEventReaderEventCommittee) ReadCreateEventCommittee(ctx context.Context, topic kafkatopic.EventKafkaEventPublisher, handler func(ctx context.Context, payload CreateEventCommitteeRequestPayload) error) error {
	config := NewKafkaReaderConfig(kc.groupID, string(topic))
	reader := kafka.NewReader(config)

	defer reader.Close()

	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		log.Log.Errorf(ctx, "[CreateEventCommittee (event reader), reader.ReadMessage (kafka)] with error detail %s", err.Error())
		return err
	}

	var data CreateEventCommitteeRequestPayload
	if err := json.Unmarshal(msg.Value, &data); err != nil {
		log.Log.Errorf(ctx, "[CreateEventCommittee (event reader), json.Unmarshal] with error detail %s", err.Error())
		return err
	}

	if err := handler(ctx, data); err != nil {
		log.Log.Errorf(ctx, "[CreateEventCommittee (event reader), handler (service)] with error detail %s", err.Error())
		return err
	}

	return err
}
