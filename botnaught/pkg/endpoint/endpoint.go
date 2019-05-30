package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "go-poker-project/Botnaught/botnaught/pkg/service"
	game "github.com/gSchool/golang-curriculum-c-6/server/pkg/game"
)

// HealthRequest collects the request parameters for the Health method.
type HealthRequest struct{}

// HealthResponse collects the response parameters for the Health method.
type HealthResponse struct {
	Err error `json:"err"`
}

// MakeHealthEndpoint returns an endpoint that invokes Health on the service.
func MakeHealthEndpoint(s service.BotnaughtService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := s.Health(ctx)
		return HealthResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r HealthResponse) Failed() error {
	return r.Err
}

// ActionRequest collects the request parameters for the Action method.
type ActionRequest struct {
	Game game.Game `json:"game"`
}

// ActionResponse collects the response parameters for the Action method.
type ActionResponse struct {
	Action game.Action `json:"action"`
	Err    error       `json:"err"`
}

// MakeActionEndpoint returns an endpoint that invokes Action on the service.
func MakeActionEndpoint(s service.BotnaughtService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ActionRequest)
		action, err := s.Action(ctx, req.Game)
		return ActionResponse{
			Action: action,
			Err:    err,
		}, nil
	}
}

// Failed implements Failer.
func (r ActionResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Health implements Service. Primarily useful in a client.
func (e Endpoints) Health(ctx context.Context) (err error) {
	request := HealthRequest{}
	response, err := e.HealthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HealthResponse).Err
}

// Action implements Service. Primarily useful in a client.
func (e Endpoints) Action(ctx context.Context, game game.Game) (action game.Action, err error) {
	request := ActionRequest{Game: game}
	response, err := e.ActionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ActionResponse).Action, response.(ActionResponse).Err
}
