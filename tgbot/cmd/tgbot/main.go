package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tgbot/internal/app"
	"tgbot/pkg/logx"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	a, err := app.New(ctx)
	if err != nil {
		logx.Fatal("app initialization error", "error", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		cancel()
	}()

	a.Run(ctx)
}
