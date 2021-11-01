package types

import (
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() (b []byte, e error) {

	return
}

func (t Time) UnmarshalJSON(b []byte) error {

	return nil
}