package entity

import "time"

type Rencana struct {
	IdRencana    int
	IdUser       int
	NamaRencana  string
	WaktuRencana string
	UpdateAt     string
	DeleteAt     time.Time
}