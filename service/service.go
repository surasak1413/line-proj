package service

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type Action interface {
	webhookAction
	lineLoginAction
}

type service struct {
	ctx echo.Context
}

type Option func(sv *service) error

func New(ctx echo.Context, options ...Option) (Action, error) {
	sv := new(service)

	for _, opt := range options {
		if err := opt(sv); err != nil {
			return nil, err
		}
	}

	if err := sv.Validate(); err != nil {
		return nil, err
	}

	if ctx != nil {
		sv.ctx = ctx
	} else {
		return nil, errors.New("echo context not found")
	}

	return sv, nil
}

func (sv *service) Validate() error {
	return nil
}
