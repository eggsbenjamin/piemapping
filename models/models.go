package models

type Journey struct {
	Id                 string `json:"id"`
	Departure_location string `json:"departure_location"`
	Arrival_location   string `json:"arrival_location"`
	Departure_time     string `json:"departure_time"`
	Arrival_time       string `json:"arrival_time"`
}
