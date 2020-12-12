package model

import (
	"errors"
	"fmt"
)

var (
	ErrLeagueNotFound = errors.New("league not found")
)

// InvalidParameterError is returned when a parameter is invalid.
type InvalidParameterError struct {
	Parameter string
}

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter: %s", i.Parameter)
}
