package types 

import "fmt"

type Float float64

func (n Float) MarshalJSON() (b []byte, e error) {

	return
}

func (n Float) UnmarshalJSON(b []byte) error {
	n = 10
	return nil
}