package service

import (
	"context"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	service, ctx := setup()

	s, err := service.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	ok := s == "ok"
	if !ok {
		t.Errorf("expected service to be ok")
	}
}

func TestGet(t *testing.T) {
	service, ctx := setup()
	d, err := service.Get(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	time := time.Now()
	today := time.Format(layout)

	// testing today's date
	ok := today == d
	if !ok {
		t.Errorf("expected dates to be equal")
	}
}

func TestValidate(t *testing.T) {
	service, ctx := setup()
	valid_date, valid_err := service.Validate(ctx, "2020-04-01 15:04:05")

	if valid_err != nil {
		t.Errorf("Error: %s", valid_err)
	}

	// testing that the date is valid
	if !valid_date {
		t.Errorf("date should be valid")
	}

	// testing an invalid date
	invalid_date, invalid_err := service.Validate(ctx, "13/01/2020")

	if invalid_err != nil {
		t.Errorf("Error: %s", invalid_err)
	}

	if invalid_date {
		t.Errorf("date should be invalid")
	}
}

func setup() (service Service, ctx context.Context) {
	return NewService(), context.Background()
}
