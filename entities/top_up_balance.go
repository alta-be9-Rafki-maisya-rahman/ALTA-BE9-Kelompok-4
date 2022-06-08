package entities

import "time"

type TopUpBalance struct {
	ID           int
	Telp         string
	UserId       int
	NominalTopUp int
	Tanggal      time.Time
}
