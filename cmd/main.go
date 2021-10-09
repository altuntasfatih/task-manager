package main

import (
	"context"
	_ "github.com/altuntasfatih/task-manager/docs"
	"github.com/altuntasfatih/task-manager/internal/app"
)

// @title Task Manager
// @version 1.0
// @Schemes  http
// @BasePath /
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app.Run(ctx, cancel)
}
