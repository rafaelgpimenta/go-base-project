package main

import (
	"context"
	"fmt"

	"resource-management/internal/application"
)

func main() {
	controller := application.NewController(nil)

	count, err := controller.GetResourcesCount(context.Background())
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}
	fmt.Printf("You have %d resources\n", count)
}
