package metrics

import "github.com/prometheus/client_golang/prometheus"

// nolint:gochecknoglobals
var (
	errorCounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   "wajve",
		Subsystem:   "",
		Name:        "error_total",
		Help:        "Count of errors",
		ConstLabels: nil,
	}, []string{"msg"})
	httpRequestDurHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "wajve",
		Subsystem:   "",
		Name:        "http_request_duration_seconds",
		Help:        "Duration of HTTP request",
		ConstLabels: nil,
		Buckets:     prometheus.DefBuckets,
	}, []string{"path", "code", "method"})
)

func Init() {
	prometheus.MustRegister(errorCounterVec)
	prometheus.MustRegister(httpRequestDurHistogram)
}

func Error(msg string) {
	labels := prometheus.Labels{
		"msg": msg,
	}

	errorCounterVec.With(labels).Inc()
}
