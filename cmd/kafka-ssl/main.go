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
	// 	logger.ErrorCtx(ctx, "Error reading CA certificate", err)
	// 	panic("Error reading CA certificate: " + err.Error())
	// }

	// caPool := x509.NewCertPool()
	// if !caPool.AppendCertsFromPEM(caCert) {
	// 	logger.ErrorCtx(ctx, "Fail to append CA certificate to pool", nil)
	// 	panic("Fail to append CA certificate to pool")
	// }

	ADDRESS := "localhost:9096"
	TOPIC := "test-ssl"
	traceId := uuid.New().String()
	ctx := context.WithValue(context.Background(), logger.TraceKey, traceId)

	tlsConfig := &tls.Config{
		RootCAs: nil, // or caPool
	}

	// === Config SASL/SCRAM with admin user ===
	mechanism, err := scram.Mechanism(scram.SHA512, "admin", "admin-secret")
	if err != nil {
		logger.ErrorCtx(ctx, "Error creating SCRAM mechanism", err)
		panic("Error creating SCRAM mechanism: " + err.Error())
	}

	// === topic creation via auto.create ===
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		TLS:           tlsConfig,
		SASLMechanism: mechanism,
	}
	conn, err := dialer.DialLeader(ctx, "tcp", ADDRESS, TOPIC, 0)
	if err != nil {
		logger.ErrorCtx(ctx, "Error creating/opening connection to topic", err)
		panic("Error creating/opening connection to topic: " + err.Error())
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

	logger.InfoCtx(ctx, "Producing message in SSL")
	err = producer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("key1"),
			Value: []byte("Hello secure Kafka!"),
		},
	)
	if err != nil {
		logger.ErrorCtx(ctx, "Error producing message", err)
		panic("Error producing message: " + err.Error())
	}

	// === Consumer ===
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{ADDRESS},
		Topic:   TOPIC,
		GroupID: "test-group",
		Dialer:  dialer,
	})
	defer reader.Close()

	logger.InfoCtx(ctx, "Consuming messages in SSL")
	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		logger.ErrorCtx(ctx, "Error consuming message", err)
		panic("Error consuming message: " + err.Error())
	}

	logger.InfoCtx(ctx, "Received message in SSL",
		"key", string(msg.Key), "value", string(msg.Value))
}
