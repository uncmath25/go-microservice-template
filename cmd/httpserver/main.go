/*
HTTP / JSON server for testing the process_name service
*/
package main

import (
	"io"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/uncmath25/go-microservice-template/networking"
	"github.com/uncmath25/go-microservice-template/service"
)

const (
	localUrl = "localhost:8080"
	logFile  = "temp.log"
)

var (
	loggingOptions = []level.Option{
		level.AllowError(),
		level.AllowInfo(),
		level.AllowDebug(),
	}
	lh      loggerHandler
	handler http.Handler
)

func init() {
	lh = loggerHandler{logFilePath: logFile}
	lh.init()

	processNameService := service.NewService(lh.logger)

	handler = networking.MakeHTTPHandler(processNameService, lh.logger)
}

func main() {
	defer lh.close()
	lh.logger.Log(http.ListenAndServe(localUrl, handler))
}

type loggerHandler struct {
	logFilePath   string
	logFileHandle *os.File
	logger        log.Logger
}

func (lh *loggerHandler) init() {
	var err error
	lh.logFileHandle, err = os.Create(lh.logFilePath)
	if err != nil {
		panic(err)
	}

	lh.logger = log.NewLogfmtLogger(log.NewSyncWriter(io.MultiWriter(os.Stdout, lh.logFileHandle)))
	lh.logger = level.NewFilter(lh.logger, loggingOptions...)
	lh.logger = log.With(lh.logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	level.Info(lh.logger).Log("event", "Starting test server")
}

func (lh *loggerHandler) close() {
	level.Info(lh.logger).Log("event", "Terminating test server")
	lh.logFileHandle.Close()
}
