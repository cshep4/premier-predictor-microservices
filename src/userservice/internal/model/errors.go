package model

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// InvalidParameterError is returned when a required parameter passed to New is invalid.
type InvalidParameterError struct {
	Parameter string
}

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter: %s", i.Parameter)
}
