package routesv1

import (
	"github.com/buaazp/fasthttprouter"
)

func InitRoutes() *fasthttprouter.Router {
	router := fasthttprouter.New()

	HealthCheck(router)
	PrometheusMetrics(router)

	return router
}
