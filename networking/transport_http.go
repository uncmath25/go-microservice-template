package networking

import (
	"context"
	mainhttp "net/http"
	"net/url"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/uncmath25/go-microservice-template/service"
)

func MakeHTTPHandler(processNameService service.Service, logger log.Logger) mainhttp.Handler {
	router := mux.NewRouter()
	e := MakeServerEndpoints(processNameService, logger)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(buildLoggingEncodeGenericError(logger)),
	}

	router.Methods("GET").Path(endpointsRoutes["processname"]).Handler(
		httptransport.NewServer(
			e.ProcessNameEndpoint,
			buildLoggingDecodeProcessNameRequest(logger),
			buildLoggingEncodeGenericResponse(logger),
			options...,
		))

	return router
}

func buildLoggingEncodeGenericResponse(logger log.Logger) func(_ context.Context, w mainhttp.ResponseWriter, response interface{}) error {
	return func(_ context.Context, w mainhttp.ResponseWriter, response interface{}) error {
		level.Info(logger).Log("event", "Started encoding response")
		defer level.Info(logger).Log("event", "Returning encoded generic response")
		return EncodeResponse(w, mainhttp.StatusOK, response)
	}
}

func buildLoggingEncodeGenericError(logger log.Logger) func(_ context.Context, err error, w mainhttp.ResponseWriter) {
	return func(_ context.Context, err error, w mainhttp.ResponseWriter) {
		level.Error(logger).Log("err", "Returning error response")
		if err == nil {
			panic("encodeError with nil error")
		}
		EncodeErrorResponse(w, mainhttp.StatusNotFound, err)
	}
}

func buildLoggingDecodeProcessNameRequest(logger log.Logger) func(_ context.Context, r *mainhttp.Request) (interface{}, error) {
	return func(_ context.Context, r *mainhttp.Request) (interface{}, error) {
		level.Info(logger).Log("event", "Decoding processname request")
		defer level.Info(logger).Log("event", "Successfully decoded processname request")

		level.Debug(logger).Log("request_url", r.URL)
		for key, val := range r.Header {
			headerVal, _ := url.Parse(val[0])
			level.Debug(logger).Log("request_header_"+key, headerVal)
		}

		vars := mux.Vars(r)
		name := vars["name"]

		// u := r.URL
		// paramMap, _ := url.ParseQuery(u.RawQuery)
		//
		// var name string
		// if val, ok := paramMap["name"]; ok {
		// 	name = val[0]
		// } else {
		// 	level.Error(logger).Log("err", errBadParams.Error())
		// 	return nil, errBadParams
		// }

		req := processNameRequest{Name: name}
		return req, nil
	}
}
