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
	conn, err := dialer.DialLeader(ctx, "tcp", ADDRESS, TOPIC, 0)
	if err != nil {
		logger.ErrorContext(ctx, "Error creating/opening connection to topic", err)
		panic("Error creating/opening connection to topic: " + err.Error())
	}
	_ = conn.Close()

	// === Producer PLAINTEXT ===
	producer := kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     TOPIC,
		BatchSize: 1,
	}
	defer producer.Close()

	logger.InfoContext(ctx, "Producing message in PLAINTEXT")
	err = producer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("key"),
			Value: []byte("Hello Kafka PLAINTEXT!"),
		},
	)
	if err != nil {
		logger.ErrorContext(ctx, "Error producing message", err)
		panic("Error  producing message: " + err.Error())
	}

	// === Consumer PLAINTEXT ===
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   TOPIC,
		GroupID: "plaintext-group",
	})
	defer reader.Close()

	logger.InfoContext(ctx, "Consuming messages in PLAINTEXT")
	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		logger.ErrorContext(ctx, "Error consuming message", err)
		panic("Error consuming message: " + err.Error())
	}

	logger.InfoContext(ctx, "Received message in PLAINTEXT",
		"key", string(msg.Key), "value", string(msg.Value))
}
