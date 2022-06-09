package entities

import "time"

type Transfer struct {
	ID               int
	NominalTransfer  int
	TransferReceiver int
	TransferUser     int
	Tanggal          time.Time
}
