package main

import (
	"fmt"

	"github.com/ropel12/Latihan/db"
	"github.com/ropel12/Latihan/db/migration"
)

func main() {
	db := db.InitDb()
	err := migration.Migration(db)
	if err != nil {
		fmt.Println(err)
	}

}
