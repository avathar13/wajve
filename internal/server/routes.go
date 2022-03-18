package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"wajve/internal/metrics"
)

func (h *Handler) initRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(metrics.NewMiddleware().Handle)    // add middleware to measure http requests duration and count
	router.Handle("/metrics", promhttp.Handler()) // host for prometheus metrics

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/trivia", h.Get).Methods(http.MethodGet)
	api.HandleFunc("/trivia/populate", h.Populate).Methods(http.MethodPost)

	return router
}
