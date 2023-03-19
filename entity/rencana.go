package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Rencana struct {
	IdRencana    int
	IdUser       int
	NamaRencana  string
	WaktuRencana time.Time
	UpdateAt     mysql.NullTime
	DeleteAt     mysql.NullTime
}
