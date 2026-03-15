package main

import (
	"context"
	"resource-management/internal/lib/logger"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func main() {

	ADDRESS := "localhost:9092"
	TOPIC := "test-plaintext"
	traceId := uuid.New().String()
	ctx := context.WithValue(context.Background(), "traceId", traceId)

	// === topic creation via auto.create ===
	dialer := &kafka.Dialer{
		Timeout: 10 * time.Second,
	}
	conn, err := dialer.DialLeader(context.Background(), "tcp", ADDRESS, TOPIC, 0)
	if err != nil {
		logger.Panic().Ctx(ctx).Stack().Err(err).
			Msg("Error creating/opening connection to topic")
	}
	_ = conn.Close()

	// === Producer PLAINTEXT ===
	producer := kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     TOPIC,
		BatchSize: 1,
	}
	defer producer.Close()

	logger.Info().Ctx(ctx).Msg("Producing message in PLAINTEXT")
	err = producer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key"),
			Value: []byte("Hello Kafka PLAINTEXT!"),
		},
	)
	if err != nil {
		logger.Panic().Ctx(ctx).Stack().Err(err).
			Msg("Error producing message")
	}

	// === Consumer PLAINTEXT ===
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   TOPIC,
		GroupID: "plaintext-group",
	})
	defer reader.Close()

	logger.Info().Ctx(ctx).Msg("Consuming messages in PLAINTEXT")
	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		logger.Panic().Ctx(ctx).Stack().Err(err).
			Msg("Error consuming message")
	}

	logger.Info().Ctx(ctx).Str("key", string(msg.Key)).
		Str("value", string(msg.Value)).Msg("Received message in PLAINTEXT")
}
