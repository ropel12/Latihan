package Testing

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ropel12/Latihan/helper"
)

func InitDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/latihan?parseTime=true")
	helper.PanicIfError(err)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
