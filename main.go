package main

import (
	"fmt"
	"log"
	routesv1 "promgo/internal/v1/routes"
	"promgo/pkg/middlewares.go"

	"github.com/valyala/fasthttp"
)

func main() {
	initFastHttp()
}

const defaultPort = "3000"

func initFastHttp() fasthttp.RequestHandler {
	var fasthttpHandler fasthttp.RequestHandler

	router := routesv1.InitRoutes()

	handlerwithLogging := middlewares.LoggerMiddleware(router.Handler)

	log.Println("Starting http server on port: ", defaultPort)
	server := fasthttp.Server{
		Name:               "hrhelpers-controller",
		ReadBufferSize:     4096 * 3,
		Handler:            handlerwithLogging,
		MaxRequestBodySize: 100560416,
	}
	log.Println("Starting http server on port: ", defaultPort)
	if err := server.ListenAndServe(":" + defaultPort); err != nil {
		//Throw
		fmt.Errorf("Error starting http server: %v", err)
	}
	return fasthttpHandler
}
