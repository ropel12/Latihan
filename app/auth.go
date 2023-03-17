package app

import (
	"fmt"
	"time"

	"github.com/ropel12/Latihan/entity"
	"github.com/ropel12/Latihan/helper"
)

func (app *App) LoginForm() {
	fmt.Print("\x1bc")
	var username, password, Repeatlogin string
	fmt.Println("\n==========================Login Form================================")
	fmt.Printf("Username : ")
	fmt.Scanln(&username)
	fmt.Printf("Password : ")
	fmt.Scanln(&password)
	data, err := app.usersRepo.FindByUsername(username)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Print("Loggin Again?(y/t): ")
		fmt.Scanln(&Repeatlogin)
		if Repeatlogin == "y" {
			app.LoginForm()
		}
		return

	}
	if data.Password != password {
		fmt.Println("Password Anda Salah")
		fmt.Print("Login Lagi? (y/t): ")
		fmt.Scanln(&Repeatlogin)
		if Repeatlogin == "y" {
			app.LoginForm()
		}
		return

	}
	if data.StatusAkun != 1 {
		fmt.Println("Akun Anda Sudah Tidak Aktif")
		fmt.Print("Login Lagi? (y/t): ")
		fmt.Scanln(&Repeatlogin)
		if Repeatlogin == "y" {
			app.LoginForm()
		}
		return
	}
	app.Session[data.Username] = data
	app.DasboardUser()
}

func (app *App) Logout() {
	key := helper.GetUser(app.Session)
	delete(app.Session, key)
}

func (app *App) Register() {
	var username, password string

	fmt.Println("================== Login Form =========================")
	fmt.Println()
	fmt.Print("Masukan Username Anda")
	fmt.Scanln(&username)
	fmt.Print("\n Masukan Password Anda")
	fmt.Scanln(&password)
	_, err := app.usersRepo.FindByUsername(username)
	if err == nil {
		app.usersRepo.CreateUser(entity.User{Username: username, Password: password})
		fmt.Println("Sukses Membuat Akun dan akan diarahkan ke menu login dalam 3 detik")
		time.Sleep(3 * time.Second)
		app.LoginForm()
	}
	fmt.Println("Akun Telah Terdaftar Silahkan Gunakan Username Yang berbeda")
	app.Register()
}
