package main

import (
	"go.uber.org/zap"

	"wajve/internal/configs"
	"wajve/internal/logger"
	"wajve/internal/metrics"
	"wajve/internal/server"
	"wajve/internal/services/samples"
)

func main() {
	log := logger.Init()
	metrics.Init()

	cfg, err := configs.Init()
	if err != nil {
		log.Fatal("cannot initialize configs", zap.Error(err))
	}

	log.Info("config is initialized")

	samplesService, err := samples.New(cfg)
	if err != nil {
		log.Fatal("cannot initialize samplesService", zap.Error(err))
	}

	log.Info("samplesService is initialized")

	h := server.New(cfg, log, samplesService)
	h.RunHTTPServer()
}
