/*
HTTP / JSON server for testing the process_name service
*/
package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/uncmath25/go-microservice-template/networking"
	"github.com/uncmath25/go-microservice-template/service"
)

var (
	loggingOptions = []level.Option{
		level.AllowError(),
		level.AllowInfo(),
		level.AllowDebug(),
	}
	logger  log.Logger
	handler *networking.LambdaHandler
)

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, loggingOptions...)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	processNameService := service.NewService(logger)

	handler = networking.MakeLambdaHandler(processNameService, logger)
}

func main() {
	defer level.Info(logger).Log("event", "Terminating test server")
	lambda.Start(handler.Handle)
}
