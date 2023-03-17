package main

import (
	"fmt"

	"github.com/ropel12/Latihan/db"
	// "github.com/ropel12/Latihan/db/migration"
	"github.com/ropel12/Latihan/repository"
)

func main() {
	db := db.InitDb()
	rencana := repository.InitRencanaRepo(db) 
	res2, err := rencana.FindRencanaByUID(3)
	res, _ := rencana.FindByRencanaId(1)
	fmt.Println(res)
	fmt.Println(res2, err)
}
