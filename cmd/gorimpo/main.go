package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lmittmann/tint"
)

var Version = "dev"

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.TimeOnly,
		AddSource:  true,
	}))
	slog.SetDefault(logger)

	logger.Info("gorimpo started!", slog.String("version", Version))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Info("booting")

	go func() {
		logger.Info("starting search loop")

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				logger.Info("search loop stopped by system")
				return
			case t := <-ticker.C:
				logger.Debug("searching...", slog.Time("tick", t))
			}
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	sig := <-stopChan
	logger.Warn("graceful shutdown...", slog.String("signal", sig.String()))
	logger.Info("shutdown notification")

	cancel()
	time.Sleep(2 * time.Second)

	logger.Info("👋 gorimpo stopped succesfuly. bye!")
}
