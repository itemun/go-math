package tester

import (
	"errors"
)

var (
	ErrTestIsOver     = errors.New("test is over")
	ErrTestNotStarted = errors.New("test is not started")
	ErrWrongAnswer    = errors.New("wrong answer")
)
