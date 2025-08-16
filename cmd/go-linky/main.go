package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"rfd59/go-linky/cmd/go-linky/core"
	"rfd59/go-linky/cmd/go-linky/services"
	"syscall"
	"time"
)

func main() {
	if err := startup(); err != nil {
		slog.Error(err.Error())
	}

	slog.Info("Go-Linky stopped.")
}

func startup() error {
	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel()}))
	slog.SetDefault(logger)

	// Channel to exit properly Go-Linky
	quit := exitChannel()

	// Load dependency injection
	settingsService, err := services.NewSettingsService(&services.SerialService{})
	if err != nil {
		return err
	}
	mqttService := services.NewMqttService(settingsService.Get().Mqtt)
	linkyService := &services.LinkyService{}

	slog.Info("Go-Linky started.", "settings", settingsService.Get())
	for {
		select {
		case <-quit:
			return nil
		default:
			if err := core.Run(settingsService.Get(), linkyService, mqttService); err != nil {
				return err
			}
			slog.Debug("Sleep pause...")
			time.Sleep(10 * time.Second)
		}
	}
}

func exitChannel() chan bool {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Channel to tell main goroutine to exit
	quit := make(chan bool, 1)

	// Goroutine to handle signals
	go func() {
		sig := <-sigs
		slog.Info(fmt.Sprintf("Received signal %q. Exiting ...", sig))
		quit <- true
	}()

	return quit
}

func getLogLevel() slog.Leveler {
	if os.Getenv("GOLINKY_DEBUG") != "" {
		return slog.LevelDebug
	} else {
		return slog.LevelInfo
	}
}
