package main

import (
	"context"
	"crypto/tls"
	"resource-management/internal/lib/logger"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func main() {
	// // === Load CA certificate ===
	// caCert, err := os.ReadFile("scripts/local-dev-ca.crt")
	// if err != nil {
	//  logger.Panic().Ctx(ctx).Stack().Err(err).
	//		Msg("Error reading CA certificate")
	// }

	// caPool := x509.NewCertPool()
	// if !caPool.AppendCertsFromPEM(caCert) {
	//  logger.Panic().Ctx(ctx).Stack().Err(err).
	//		Msg("Fail to append CA certificate to pool")
	// }

	ADDRESS := "localhost:9096"
	TOPIC := "test-ssl"
	traceId := uuid.New().String()
	ctx := context.WithValue(context.Background(), "traceId", traceId)

	tlsConfig := &tls.Config{
		RootCAs: nil, // or caPool
	}

	// === Config SASL/SCRAM with admin user ===
	mechanism, err := scram.Mechanism(scram.SHA512, "admin", "admin-secret")
	if err != nil {
		logger.Panic().Ctx(ctx).Stack().Err(err).
			Msg("Error creating SCRAM mechanism")
	}

	// === topic creation via auto.create ===
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		TLS:           tlsConfig,
		SASLMechanism: mechanism,
	}
	conn, err := dialer.DialLeader(context.Background(), "tcp", ADDRESS, TOPIC, 0)
	if err != nil {
		logger.Panic().Ctx(ctx).Stack().Err(err).
			Msg("Error creating/opening connection to topic")
	}
	_ = conn.Close()

	// === Producer ===
	producer := kafka.Writer{
		Addr:      kafka.TCP(ADDRESS),
		Topic:     TOPIC,
		BatchSize: 1,
		Transport: &kafka.Transport{
			TLS:  tlsConfig,
			SASL: mechanism,
		},
	}
	defer producer.Close()

	logger.Info().Ctx(ctx).Msg("Producing message in SSL")
	err = producer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key1"),
			Value: []byte("Hello secure Kafka!"),
		},
	)
	if err != nil {
		logger.Panic().Ctx(ctx).Stack().Err(err).
			Msg("Error producing message")
	}

	// === Consumer ===
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{ADDRESS},
		Topic:   TOPIC,
		GroupID: "test-group",
		Dialer:  dialer,
	})
	defer reader.Close()

	logger.Info().Ctx(ctx).Msg("Consuming messages in SSL")
	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		logger.Panic().Ctx(ctx).Stack().Err(err).
			Msg("Error consuming message")
	}

	logger.Info().Ctx(ctx).Str("key", string(msg.Key)).
		Str("value", string(msg.Value)).Msg("Received message in SSL")
}
