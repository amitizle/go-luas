package luas

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestXMLUnmarshalling(t *testing.T) {
	testTable := []struct {
		Name        string
		TestXMLFile string
		Stop        string
		StopAbv     string
		Message     string
		Inbound     Direction
		Outbound    Direction
	}{
		{
			"Balally",
			"balally.xml",
			"Balally",
			"BAL",
			"Green Line services operating normally",
			Direction{
				Name: "Inbound",
				Trams: []Tram{
					Tram{
						DueMins:     "3",
						Destination: "Parnell",
					},
					Tram{
						DueMins:     "DUE",
						Destination: "Broombridge",
					},
				},
			},
			Direction{
				Name: "Outbound",
				Trams: []Tram{
					Tram{
						DueMins:     "9",
						Destination: "Bride's Glen",
					},
				},
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			xmlDataFile := filepath.Join("testdata", testCase.TestXMLFile)
			bytes, err := ioutil.ReadFile(xmlDataFile)
			if err != nil {
				t.Fatalf("error reading fixtures file %v, %v", testCase.TestXMLFile, err)
			}
			stopInfo, err := parseLuasResponse(bytes)
			if err != nil {
				t.Fatalf("cannot parse luas response xml, %v", err)
			}
			assert(t, "stop should be equal", testCase.Stop, stopInfo.Stop)
			assert(t, "stop abv should be equal", testCase.StopAbv, stopInfo.StopAbv)
			assert(t, "message should be equal", testCase.Message, stopInfo.Message)
			// TODO test tram by tram
		})
	}
}

func TestGetStop(t *testing.T) {
	testTable := []struct {
		Name         string
		StopName     string
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
		{"get existing stop by long name", "Balally", &Stop{
			Name:        "Balally",
			NameAbv:     "BAL",
			Line:        "green",
			Coordinates: []float64{53.28610556, -6.236772222},
		}, nil},
	}

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			stop, err := GetStop(testCase.StopName)
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

func assert(t *testing.T, message string, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf(message, "%s: expected %v, got %v", message, a, b)
	}
}
