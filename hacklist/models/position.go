package models

import ()

type Position struct {
	Country     string
	Province    string
	City        string
	Section     string
	Ip          string
	TeleCompany string
	Geo         []float64
}

type ptn struct{}

var Ptn ptn

func (ptn) New() {}
