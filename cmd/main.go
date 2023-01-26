package main

import (
	"context"

	_ "github.com/altuntasfatih/car-service-backend/docs"
	"github.com/altuntasfatih/car-service-backend/internal/app"
)

// @title Car Service
// @version 1.0
// @Schemes  http
// @BasePath /
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app.Run(ctx, cancel)
}
