package service

import (
	"context"
	"encoding/json"
	"net/http"
)

// getRequest is used to format the JSON payload of a request
type getRequest struct{}

// getResponse is used to format the JSON payload of a response
type getResponse struct {
	Date string `json:"date"`
	Err  string `json:"err,omitempty"`
}

// validateRequest to format JSON of a valid request
type validateRequest struct {
	Date string `json:"date"`
}

// validateResponse to format JSON of a valid response
type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

// statusRequest is the format for requesting the status
type statusRequest struct{}

// statusResponse is the format for response of the status
type statusResponse struct {
	Status string `json:"status"`
}

// decodeGetRequest decodes a get request
func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	return req, nil
}

// decodeValidateRequest decodes a valid request
func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// decodeStatusRequest decodes status request
func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

// encodeResponse the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
