package app

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bakhram74/wallet-app-test-work/config"
	"github.com/Bakhram74/wallet-app-test-work/internal/controller/http"
	"github.com/Bakhram74/wallet-app-test-work/internal/repository"
	"github.com/Bakhram74/wallet-app-test-work/internal/service"
	"github.com/Bakhram74/wallet-app-test-work/pkg/httpserver"
	"github.com/Bakhram74/wallet-app-test-work/pkg/postgres"
)

func Run(cfg *config.Config) {
	
	slog.Debug("Postgresql initializing")
	slog.Debug("maximum size of the pool", slog.Int("PG_POOL_MAX", cfg.PoolMax))

	pg, err := postgres.New(cfg.PGUrl, postgres.MaxPoolSize(cfg.PoolMax))
	if err != nil {
		panic(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	err = RunMigration(cfg)
	if err != nil {
		panic(fmt.Sprintf("Migration error: %s", err.Error()))
	}

	repo := repository.New(pg.Pool)

	service := service.NewService(repo)

	handler := http.NewHandler(cfg, service).Init()

	httpServer := httpserver.New(handler, httpserver.Port(cfg.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		slog.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	case err := <-httpServer.Notify():
		slog.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err).Error())
	}

	err = httpServer.Shutdown()
	if err != nil {
		slog.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}

}
