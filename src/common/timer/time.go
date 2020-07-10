package timer

import "time"

type Time interface {
	Now() time.Time
}

type t struct {

}

func NewTime() Time {
	return t{}
}

func (t) Now() time.Time {
	return time.Now()
}
