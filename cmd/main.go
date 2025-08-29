package main

import (
	"context"

	"resource-management/internal/application"
	"resource-management/internal/lib/logger"

	"github.com/google/uuid"
)

func main() {
	controller := application.NewController(nil)
	traceId := uuid.New().String()
	ctx := context.WithValue(context.Background(), "traceId", traceId)

	count, err := controller.GetResourcesCount(ctx)
	if err != nil {
		logger.Error().Ctx(ctx).Msg("Something went wrong when getting resources")
		return
	}
	logger.Info().Ctx(ctx).Int32("count", count).Msg("Successfully found resources")
}
