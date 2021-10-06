package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func interruptToCancel(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		signalStream := make(chan os.Signal)
		defer func() {
			signal.Stop(signalStream)
			close(signalStream)
		}()
		signal.Notify(signalStream, os.Interrupt, syscall.SIGTERM)
		select {
		case <-signalStream:
			cancel()
			return
		case <-ctx.Done():
			return
		}
	}()
}

func Run(ctx context.Context, cancel context.CancelFunc) {
	var wg sync.WaitGroup
	interruptToCancel(ctx, cancel, &wg)

	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}
	app.Listen(ctx, &wg)

	wg.Wait()
}



