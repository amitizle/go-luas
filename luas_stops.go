package luas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var configFile = "luas_stops.json"

type Stop struct {
	Name        string    `json:"name"`
	NameAbv     string    `json:"name_abv"`
	Line        string    `json:"line"`
	Coordinates []float64 `json:"coordinates"`
}

func AllStops() ([]Stop, error) {
	stops := make([]Stop, 0)
	stopsBytes, err := ioutil.ReadFile(filepath.Join("configs", configFile))
	if err != nil {
		return stops, nil
	}
	err = json.Unmarshal(stopsBytes, &stops)
	if err != nil {
		return stops, nil
	}
	return stops, nil
}

func GetStop(stopAbv string) (*Stop, error) {
	stopAbvUpcase := strings.ToUpper(stopAbv)
	allStops, err := AllStops()
	if err != nil {
		return &Stop{}, err
	}
	for _, stop := range allStops {
		if stop.NameAbv == stopAbvUpcase {
			return &stop, nil
		}
	}
	return &Stop{}, fmt.Errorf("no such stop %v", stopAbvUpcase)
}
