package entities

import "time"

type UserBalance struct {
	IdUser  int
	Saldo   int
	Tanggal time.Time
}
