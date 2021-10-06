package main

import (
	"context"
	"github.com/altuntasfatih/task-manager/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app.Run(ctx, cancel)
}