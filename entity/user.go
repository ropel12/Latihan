package entity

import "github.com/go-sql-driver/mysql"

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
	WaktuRencana string
	UpdateAt     mysql.NullTime
	DeleteAt     mysql.NullTime
}
