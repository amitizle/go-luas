package luas

import (
	"encoding/xml"
	"fmt"
	"github.com/amitizle/go-luas/internal/http_client"
)

type StopInfo struct {
	Directions []Direction `xml:"direction"`
	Message    string      `xml:"message"`
	Created    string      `xml:"created,attr"`
	Stop       string      `xml:"stop,attr"`
	StopAbv    string      `xml:"stopAbv,attr"` // Short name for the stop (3 letters)
}

type Direction struct {
	Name  string `xml:"name,attr"` // Inbound/Outbound
	Trams []Tram `xml:"tram"`
}

type Tram struct {
	DueMins     int    `xml:"dueMins,attr"`
	Destination string `xml:"destination,attr"`
}

type Stop struct {
	NameAbv     string
	Name        string
	Line        string
	Coordinates map[string]float64
}

func GetStop(stopAbv string) (*Stop, error) {
}

// func GetStopInfo(stopAbv string) (*StopInfo, error) {
// 	httpC, _ := luas_http_client.NewClient("")
// 	resp, err := httpC.GetForecast(stopAbv)
// 	if err != nil {
// 		return &StopInfo{}, err
// 	}
// 	stopInfo, err := parseLuasResponse(resp)
// 	if err != nil {
// 		return &StopInfo{}, err
// 	}
// 	return stopInfo, nil
// }

func main() {
	httpC, _ := luas_http_client.NewClient("")
	resp, _ := httpC.GetForecast()

	fmt.Println("Body:", string(resp))

	var q StopInfo
	xml.Unmarshal(resp, &q)
	fmt.Println(q)

}

func parseLuasResponse(xmlString []byte) (*StopInfo, error) {
	var stopInfo StopInfo
	err := xml.Unmarshal(xmlString, &stopInfo)
	if err != nil {
		return &StopInfo{}, err
	}
	return &stopInfo, nil
}
