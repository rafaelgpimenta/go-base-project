package main

import (
	"context"

	"resource-management/internal/application/controllers"
	"resource-management/internal/domain/interfaces"
	"resource-management/internal/lib/config"
	"resource-management/internal/lib/logger"

	"github.com/google/uuid"
)

func main() {
	controller := controllers.NewResourceController(nil)
	traceId := uuid.New().String()
	ctx := context.WithValue(context.Background(), logger.TraceKey, traceId)
	testConfig, _ := config.Get[interfaces.TestConfig]("test")

	count, err := controller.GetResourcesCount(ctx)
	if err != nil {
		logger.ErrorCtx(ctx, "Something went wrong when getting resources", err)
		return
	}
	logger.InfoCtx(ctx, "Successfully found resources",
		"count", count, "foo", testConfig.Foo, "fooInt", testConfig.FooInt,
		"fooFloat", testConfig.FooFloat, "fooBoolean", testConfig.FooBoolean)
}
