package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/Bakhram74/wallet-app-test-work/config"
	"github.com/Bakhram74/wallet-app-test-work/internal/app"
)

func ParseLevel(s string) (slog.Level, error) {
	var level slog.Level
	err := level.UnmarshalText([]byte(s))
	return level, err
}

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	level, err := ParseLevel(cfg.Level)
	if err != nil {
		log.Fatalf("Logger error: %s", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: level}))
	slog.SetDefault(logger)

	app.Run(cfg)
}
