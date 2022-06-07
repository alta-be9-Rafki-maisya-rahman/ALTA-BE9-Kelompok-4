package entities

import "time"

type TopUpBalance struct {
	ID      int
	UserId  int
	NominalTopUp int
	Tanggal time.Time
}
