package main

import (
	"github.com/alserov/auth-service/internal/app"
	"github.com/alserov/auth-service/internal/config"
	"github.com/alserov/auth-service/internal/lib/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.MustSetupLogger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	s := <-stop
	application.GRPCServer.Stop()
	log.Info("application stopped", slog.String("signal", s.String()))
}
