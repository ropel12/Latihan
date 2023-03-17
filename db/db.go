package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ropel12/Latihan/config"
	"github.com/ropel12/Latihan/helper"
)

func InitDb() *sql.DB {
	db, err := sql.Open(config.DbDriver, "root:"+config.DbPassword+"@tcp(localhost:"+config.DbPort+")/"+config.DBname+"?parseTime=true")
	helper.PanicIfError(err)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
