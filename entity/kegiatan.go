package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Kegiatan struct {
	Idkegiatan    int
	Userid        int
	NamaKegiatan  string
	WaktuKegiatan time.Time
	UpdateAt      mysql.NullTime
	DeleteAt      mysql.NullTime
}
