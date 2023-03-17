package main

import (
	"fmt"

	"github.com/ropel12/Latihan/app"
	"github.com/ropel12/Latihan/db"
	"github.com/ropel12/Latihan/db/migration"
	"github.com/ropel12/Latihan/repository"
)

func main() {

	db := db.InitDb()
	migration.Migration()
	UserRepo := repository.InitUserRepo(db)
	KegiatanRepo := repository.InitKegiatanRepo(db)
	var choice int
	defer db.Close()
	defer fmt.Println("Terimakasih telah menggunakan aplikasi kami")
	fmt.Println("Daftar Pilihan :")
	fmt.Println("1.Running Aplikasi")
	fmt.Println("9.Exit")
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&choice)
	for choice != 9 && choice == 1 {
		switch choice {
		case 1:
			app := app.NewApp(UserRepo, KegiatanRepo)
			app.Start()
		}

	}

}
