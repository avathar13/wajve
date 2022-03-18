package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Handle(next http.Handler) http.Handler {
	return m.HandleFunc(next)
}

func (m *Middleware) HandleFunc(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		handler := promhttp.InstrumentHandlerDuration(
			httpRequestDurHistogram.MustCurryWith(prometheus.Labels{
				"path": request.URL.Path,
			}),
			next,
		)

		handler(writer, request)
	}
}
