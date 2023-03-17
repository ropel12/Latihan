package entity

import "time"

type Kegiatan struct {
	Idkegiatan    int
	Userid        int
	NamaKegiatan  string
	WaktuKegiatan string
	UpdateAt      time.Time
	DeleteAt      time.Time
}
