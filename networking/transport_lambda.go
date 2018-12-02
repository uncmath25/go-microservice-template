package networking

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/uncmath25/go-microservice-template/service"
)

func MakeLambdaHandler(processNameService service.Service, logger log.Logger) *LambdaHandler {
	return &LambdaHandler{
		endpoints: MakeServerEndpoints(processNameService, logger),
		logger:    logger,
	}
}

type LambdaHandler struct {
	endpoints endpoints
	logger    log.Logger
}

func (lm *LambdaHandler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	level.Info(lm.logger).Log("event", "Decoding request")
	defer level.Info(lm.logger).Log("event", "Returning processed response")

	response, err := applyEndpoints(ctx, lm.endpoints, req)
	if err != nil {
		return returnError(err, http.StatusInternalServerError, lm.logger)
	}
	resJson, err := json.Marshal(response)
	if err != nil {
		return returnError(err, http.StatusInternalServerError, lm.logger)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(resJson),
	}, nil
}

func returnError(err error, statusCode int, logger log.Logger) (events.APIGatewayProxyResponse, error) {
	level.Error(logger).Log("err", err.Error())
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       err.Error(),
	}, nil
}

func applyEndpoints(ctx context.Context, e endpoints, req events.APIGatewayProxyRequest) (interface{}, error) {
	matchedEnpointRouteKey := ""
	for key, val := range endpointsRoutes {
		if matchParamPath(val, req.Path) {
			matchedEnpointRouteKey = key
			break
		}
	}
	if matchedEnpointRouteKey == "" {
		return nil, errBadParams
	}

	switch matchedEnpointRouteKey {
	case "processname":
		name := req.PathParameters["name"]
		return e.ProcessNameEndpoint(ctx, processNameRequest{Name: name})
	default:
		return nil, errBadParams
	}
}

func matchParamPath(route string, url string) bool {
	route_parts := strings.Split(route, "/")
	url_parts := strings.Split(url, "/")

	if len(route_parts) != len(url_parts) {
		return false
	}

	for i:=0; i<len(route_parts); i++ {
		if len(route_parts[i]) > 0 && string(route_parts[i][0]) == "{" && len(url_parts[i]) > 0 {
			continue
		}
		if route_parts[i] != url_parts[i] {
			return false
		}
	}

	return true
}
