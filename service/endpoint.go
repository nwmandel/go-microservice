package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	GetEndpoint      endpoint.Endpoint
	StatusEndpoint   endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

// MakeGetEndpoint returns the response from the get service
func MakeGetEndpoint(serv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getRequest)
		get, err := serv.Get(ctx)
		if err != nil {
			return getResponse{get, err.Error()}, nil
		}
		return getResponse{get, ""}, nil
	}
}

// MakeStatusEndpoint returns the response from the status service
func MakeStatusEndpoint(serv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusRequest)
		status, err := serv.Status(ctx)
		if err != nil {
			return statusResponse{status}, nil
		}
		return statusResponse{status}, nil
	}
}

// MakeValidateEndpoint returns the response the validate service
func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		valid, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return validateResponse{valid, err.Error()}, nil
		}
		return validateResponse{valid, ""}, nil
	}
}

// Get endpoint mapping
func (e Endpoints) Get(ctx context.Context) (string, error) {
	req := getRequest{}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	getResp := resp.(getResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Date, nil
}

// Status endpoint mapping
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := statusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	statusResp := resp.(statusResponse)
	return statusResp.Status, nil
}

// Validate endpoint mapping
func (e Endpoints) Validate(ctx context.Context, date string) (bool, error) {
	req := validateRequest{Date: date}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}

	validateResp := resp.(validateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}
