package errors

import (
	"errors"
)

// Custom errors to be used for this project
var (
	ErrUnAuthorized = errors.New("invalid username or token")
	ErrZeroInDenominator = errors.New("can't divide by 0")
)