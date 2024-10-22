package healthservices

import (
	httputils "promgo/pkg/http"

	"github.com/valyala/fasthttp"
)

func HealthCheckHandler(ctx *fasthttp.RequestCtx) {

	httputils.ResponseHandlers(ctx, nil, nil, fasthttp.StatusOK, "Health Check OK")
}
