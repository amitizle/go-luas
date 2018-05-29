package luas

import (
	"encoding/xml"
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

func GetStopForecast(stopAbv string) (*StopInfo, error) {
	stop, err := GetStop(stopAbv)
	var stopInfo StopInfo
	if err != nil {
		return &stopInfo, err
	}
	httpClient, err := luas_http_client.NewClient("")
	if err != nil {
		return &stopInfo, err
	}
	resp, err := httpClient.GetForecast(stop.NameAbv)
	if err != nil {
		return &stopInfo, err
	}
	err = xml.Unmarshal(resp, &stopInfo)

	return &stopInfo, nil
}

func parseLuasResponse(xmlString []byte) (*StopInfo, error) {
	var stopInfo StopInfo
	err := xml.Unmarshal(xmlString, &stopInfo)
	if err != nil {
		return &StopInfo{}, err
	}
	return &stopInfo, nil
}
