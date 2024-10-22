package routesv1

import (
	healthservices "promgo/internal/v1/helpers/health"

	"github.com/buaazp/fasthttprouter"
)

func PrometheusMetrics(router *fasthttprouter.Router) {
	router.GET("/health", healthservices.HealthCheckHandler)
}
