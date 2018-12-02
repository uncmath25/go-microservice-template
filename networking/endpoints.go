package networking

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/uncmath25/go-microservice-template/service"
)

var (
	errBadParams    = errors.New("Request url params were malformed")
	endpointsRoutes = map[string]string{"processname": "/process_name/{name}"}
)

func MakeServerEndpoints(processNameService service.Service, logger log.Logger) endpoints {
	return endpoints{
		ProcessNameEndpoint: makeProcessNameEndpoint(processNameService, logger),
	}
}

type endpoints struct {
	ProcessNameEndpoint endpoint.Endpoint
}

func makeProcessNameEndpoint(s service.Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		level.Info(logger).Log("event", "Passing request to processname service")
		req := request.(processNameRequest)
		level.Debug(logger).Log("name", req.Name)
		processedName, err := s.ProcessName(ctx, req.Name)
		if err != nil {
			level.Error(logger).Log("err", err.Error())
			return nil, err
		}
		level.Info(logger).Log("event", "Returning response from processname service")
		level.Debug(logger).Log("name", processedName.Name)
		return processedNameResponse{ProcessedName: *processedName}, nil
	}
}

type processNameRequest struct {
	Name string
}

type processedNameResponse struct {
	ProcessedName service.ProcessedName
}
