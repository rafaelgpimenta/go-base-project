package main

import (
	"context"

	"resource-management/internal/application"
	"resource-management/internal/domain/interfaces"
	"resource-management/internal/lib/config"
	"resource-management/internal/lib/logger"

	"github.com/google/uuid"
)

func main() {
	controller := application.NewController(nil)
	traceId := uuid.New().String()
	ctx := context.WithValue(context.Background(), "traceId", traceId)
	testConfig, _ := config.Get[interfaces.TestConfig]("test")

	count, err := controller.GetResourcesCount(ctx)
	if err != nil {
		logger.Error().Ctx(ctx).Msg("Something went wrong when getting resources")
		return
	}
	logger.Info().Ctx(ctx).Int32("count", count).
		Str("foo", testConfig.Foo).
		Int32("fooInt", testConfig.FooInt).
		Float64("fooFloat", testConfig.FooFloat).
		Bool("fooBoolean", testConfig.FooBoolean).
		Msg("Successfully found resources")
}
