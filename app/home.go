package app

import "fmt"

func (app *App) HomePage() {
	var choice int
	fmt.Println("Silahkan Pilih Menu Dibawah Ini : ")
	fmt.Println("1.Login")
	fmt.Println("2.Register")
	fmt.Println("3.Kembali")
	fmt.Print("Masukan pilihan : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		app.LoginForm()

	case 2:
		app.Register()

	case 3:
		break
	default:
		break
	}

}
