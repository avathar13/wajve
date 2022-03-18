package server

import (
	"net/http"

	"go.uber.org/zap"

	"wajve/internal/configs"
	"wajve/internal/services/samples"
)

type Handler struct {
	config         *configs.Cfg
	logger         *zap.Logger
	samplesService samples.ServiceInterface
}

func New(cfg *configs.Cfg, logger *zap.Logger, samplesService samples.ServiceInterface) *Handler {
	return &Handler{
		config:         cfg,
		logger:         logger,
		samplesService: samplesService,
	}
}

func (h *Handler) RunHTTPServer() {
	cfg := h.config
	endpoint := cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port
	router := h.initRoutes()

	server := &http.Server{
		Addr:              endpoint,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       cfg.HTTPServer.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      cfg.HTTPServer.WriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	h.logger.Info("Running server at " + server.Addr)

	if err := server.ListenAndServe(); err != nil {
		h.logger.Fatal("server.shutdown.error", zap.Error(err))
	}
}
