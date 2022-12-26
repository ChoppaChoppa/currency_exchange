package service

import "errors"

var (
	ErrBadRequest = errors.New("bad request")
	ErrServer     = errors.New("server error")
)
