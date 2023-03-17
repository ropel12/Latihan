package entity

import "time"

type User struct {
	Userid     int
	Username   string
	Password   string
	StatusAkun int
}

type Rencana struct {
	IdRencana    int
	IdUser       int
	NamaRencana  string
	WaktuRencana time.Time
	UpdateAt     time.Time
	DeleteAt     time.Time
}

