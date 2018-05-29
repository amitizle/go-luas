package luas

import (
	"fmt"
	"strings"
)

func AllStops() []*Stop {
	return allStops
}

// TODO make this a map
func GetStop(stopAbv string) (*Stop, error) {
	stopAbvUpcase := strings.ToUpper(stopAbv)
	allStops := AllStops()
	for _, stop := range allStops {
		if stop.NameAbv == stopAbvUpcase {
			return stop, nil
		}
	}
	return &Stop{}, fmt.Errorf("no such stop %v", stopAbvUpcase)
}
