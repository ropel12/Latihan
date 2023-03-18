package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

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
	UpdateAt     mysql.NullTime
	DeleteAt     mysql.NullTime
}
