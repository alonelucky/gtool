package numbers

import (
	"testing"
)

func TestToFixed(t *testing.T) {
	var (
		example1 = []float64{922.5388456553852, 203.21701043529018, 816.081289552975}
	)

	for k := range example1 {
		t.Error(ToFixed(example1[k], k+1))
	}
}

func TestRound(t *testing.T) {
	var (
		example1 = []float64{922.5388456553852, 203.21701043529018, 816.081289552975}
	)

	for k := range example1 {
		t.Error(Round(example1[k], k+1))
	}
}

func TestRound32(t *testing.T) {
	var (
		example1 = []float32{922.5388456553852, 203.21701043529018, 816.081289552975}
	)

	for k := range example1 {
		t.Error(Round32(example1[k], k+1))
	}
}

func TestRoundAll(t *testing.T) {
	var (
		example1 = []interface{}{922.5388456553852, "203.21701043529018", 816}
	)

	for k := range example1 {
		t.Error(RoundAll(example1[k], k+1))
	}
}

func TestBankRound(t *testing.T) {
	var (
		example1 = []float64{922.5388456553852, 203.21701043529018, 816.081289552975}
	)

	for k := range example1 {
		t.Error(BankRound(example1[k], k+1))
	}
}
