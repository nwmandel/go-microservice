package service

import (
	"context"
	"time"
)

// layout time format
var layout string = "2006-01-02 15:04:05"

// Service provides some "date capabilities"
type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type dateService struct{}

// NewService makes a new date service.
func NewService() Service {
	return dateService{}
}

// Status only tell us that our service is ok
func (dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get will return today's date
func (dateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format(layout), nil
}

// Validate will check if the date today's date
func (dateService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse(layout, date)
	if err != nil {
		return false, err
	}
	return true, nil
}
