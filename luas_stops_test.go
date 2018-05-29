package luas

import (
	"errors"
	"testing"
)

func TestGetStop(t *testing.T) {
	testTable := []struct {
		Name         string
		StopAbv      string
		ExpectedStop *Stop
		ExpectedErr  error
	}{
		{"existing stop should return the stop", "BAL", &Stop{
			Name:        "Balally",
			NameAbv:     "BAL",
			Line:        "green",
			Coordinates: []float64{53.28610556, -6.236772222},
		}, nil},
		{"non existing stop should return an error", "IKO", &Stop{}, errors.New("no such stop IKO")},
	}

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			stop, err := GetStop(testCase.StopAbv)
			if stop.NameAbv != testCase.ExpectedStop.NameAbv || // TODO uncomplicate this thing
				((err != nil || testCase.ExpectedErr != nil) && err.Error() != testCase.ExpectedErr.Error()) {
				t.Fatalf(
					"expected stop %v and err %v, got stop %v and err %v",
					testCase.ExpectedStop,
					testCase.ExpectedErr,
					stop,
					err,
				)
			}
		})
	}
}
