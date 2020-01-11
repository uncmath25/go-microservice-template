/*
Service which determines which show ids need to be updated for the given api
*/
package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Service interface {
	ProcessName(ctx context.Context, name string) (*ProcessedName, error)
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

type service struct {
	logger log.Logger
}

func (s *service) ProcessName(ctx context.Context, name string) (*ProcessedName, error) {
	level.Info(s.logger).Log("event", "Received request to obtain show updates")
	level.Debug(s.logger).Log("api", name)

	defer level.Info(s.logger).Log("event", "Returning show updates")
	switch name {
	case "Colton":
		return &ProcessedName{Name: "Colton", Message: "You the boss!"}, nil
	default:
		return &ProcessedName{Name: name, Message: "Youse a poser..."}, nil
	}
}
