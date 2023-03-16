package entity

import "time"

type User struct {
	Userid     int
	Username   string
	Password   string
	StatusAkun bool
}

type Kegiatan struct {
	Idkegiatan    int
	Userid        int
	NamaKegiatan  string
	WaktuKegiatan string
	UpdateAt      time.Time
	DeleteAt      time.Time
}

type Rencana struct {
	IdRencana    int
	IdUser       int
	NamaRencana  string
	WaktuRencana time.Time
	UpdateAt     time.Time
	DeleteAt     time.Time
}

type RencanaKegiatan struct {
	Idkegiatan       int
	IdRencana        int
	IdUser           int
	IDRencanKegiatan int
}
