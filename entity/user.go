package entity

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
	UpdateAt     string
	DeleteAt     string
}
