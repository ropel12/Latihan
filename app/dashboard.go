package app

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ropel12/Latihan/entity"
	"github.com/ropel12/Latihan/helper"
)

func (app *App) DasboardUser() {
	key := helper.GetUser(app.Session)
	fmt.Printf("Selamat datang %s di aplikasi kami \n", app.Session[key].Username)
	var choice int
	fmt.Println("Masukan menu yang ingin digunakan :")
	fmt.Println("1.Lihat Semua Kegiatan")
	fmt.Println("2.Lihat Semua Rencana")
	fmt.Println("3.Tambah Kegiatan")
	fmt.Println("4.Tambah Rencana")
	fmt.Println("5.Update Profile")
	fmt.Println("6.Masukan kegiatan ke rencana")
	fmt.Println("7.Logout Akun")
	fmt.Println("Masukan Pilihan Anda:")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		app.ListKegiatan()
	case 2:
	case 3:
		app.FormTambahKegiatan()
	case 4:
	case 5:
		app.UpdateProfile()
	case 6:
	case 7:
		app.Logout()
	}
}

func (app *App) UpdateProfile() {
	var choice int
	defer func() {
		fmt.Println("Jika Ingin Kembali Tekan 9")
		fmt.Scan(&choice)
		if choice == 9 {
			app.DasboardUser()
		}
		app.UpdateProfile()
	}()
	key := helper.GetUser(app.Session)
	var username, password string
	fmt.Printf("Username  Lama : %s\n", app.Session[key].Username)
	fmt.Printf("Password Baru : %s\n\n", app.Session[key].Password)
	fmt.Print("Masukan Username Baru: ")
	fmt.Scan(&username)
	_, err := app.usersRepo.FindByUsername(username)
	if err != nil {
		fmt.Println("Username Telah Terdaftar Silahkan Gunakan Yang Lain")
		app.UpdateProfile()
		return
	}
	fmt.Print("Masukan Password Baru: ")
	fmt.Scan(&password)
	err1 := app.usersRepo.UpdateUser(entity.User{Username: username, Password: password, StatusAkun: 1})
	if err1 != nil {
		fmt.Println(err1.Error())
	}

}

func (app *App) FormTambahKegiatan() {
	key := helper.GetUser(app.Session)
	var namakegiatan, choice string
	fmt.Println("=================FORM TAMBAH KEGIATAN===========================")
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukan Nama Kegiatan: ")
	fmt.Scan(&namakegiatan)
	_, err := app.KegiatanRepo.FindKegiatanByName(namakegiatan, app.Session[key].Userid)
	if err != nil {
		err2 := app.KegiatanRepo.Create(entity.Kegiatan{NamaKegiatan: namakegiatan, Userid: app.Session[key].Userid})
		if err2 != nil {
			fmt.Println(err2.Error())
			app.FormTambahKegiatan()
			return
		}
		fmt.Println("Tambah Kegiatan Berhasil Ditambahkan")
		fmt.Print("Ingin Menambahkan data lagi? (y/t)")
		fmt.Scan(&choice)
		if choice == "y" {
			app.FormTambahKegiatan()
			return
		}
		fmt.Println("Anda akan redirect ke menu utama dalam 3 detik")
		time.Sleep(time.Second * 3)
		app.DasboardUser()
		return
	}
	fmt.Println("Nama Kegiatan Sudah Ada")
	time.Sleep(time.Second * 2)
	app.FormTambahKegiatan()

}

func (app *App) ListKegiatan() {
	key := helper.GetUser(app.Session)
	var choice int
	datas, err := app.KegiatanRepo.GetAll(app.Session[key].Userid)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(time.Second * 2)
		app.DasboardUser()
	}
	fmt.Printf("Berikut ini daftar kegiatan dari %s yang telah dibuat\n", app.Session[key].Username)
	if len(datas) == 0 {
		fmt.Println("Anda belum memiliki daftar kegiatan.")
		fmt.Print("Jika Anda ingin menambahkan kegiatan masukan angka 1 Jika Tidak masukan 0: ")
		fmt.Scan(&choice)
		if choice == 1 {
			app.FormTambahKegiatan()
			return
		}
		app.DasboardUser()
		return

	}
	helper.PrintData(datas)
	fmt.Println("Jika anda ingin menambahkan lagi masukan angka 1\n jika ingin menghapus masukan angka 2 \njika ingin mengupdate masukan angka 3")
	fmt.Print("Masukan Pilihan : ")
	fmt.Scan(&choice)
	if choice == 1 {
		app.FormTambahKegiatan()
		return
	} else if choice == 2 {
		app.HapusKegiatan()
		return
	} else if choice == 3 {
		app.UpdateKegiatan()
		return
	}
	fmt.Println("Anda akan diarahkan ke halaman dashboard")
	time.Sleep(time.Second * 1)
	app.DasboardUser()
}

func (app *App) HapusKegiatan() {
	key := helper.GetUser(app.Session)
	var choices string
	datas, err := app.KegiatanRepo.GetAll(app.Session[key].Userid)
	if err != nil {
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(time.Second * 2)
			app.DasboardUser()
			return
		}
	}
	if len(datas) != 0 {
		fmt.Println("Silahkan pilih daftar kegiatan yang ingin dihapus")
		helper.PrintData(datas)
		fmt.Print("Masukan data yang ingin dihapus jika ingin banyak tambahkan koma contoh(1,2,3,4,5)")
		fmt.Scan(&choices)
		var index int
		if strings.Contains(choices, ",") {
			ids := strings.Split(choices, ",")
			for i, val := range ids {
				toint, _ := strconv.Atoi(val)
				toint -= 1
				err := app.KegiatanRepo.DeleteKegiatan(datas[toint].Idkegiatan)
				if err != nil {
					fmt.Printf("Id %d Tidak terdaftar\n", datas[toint].Idkegiatan)
					fmt.Printf("Data yang Dihapus sebanyak %d", i+1)
					break
				}
				index++
			}
			if index < 1 {
				fmt.Println("Masukan Data Yang benar")
				time.Sleep(2 * time.Second)
				app.HapusKegiatan()
				return
			}
		}
		toint, _ := strconv.Atoi(choices)
		err := app.KegiatanRepo.DeleteKegiatan(datas[toint].Idkegiatan)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Masukan Data Yang Benar")
			app.HapusKegiatan()
			return
		}
		fmt.Println("Berhasil Mengapus Data")
		fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
		fmt.Scan(&choices)
		if choices == "y" {
			app.HapusKegiatan()
			return
		}
		fmt.Println("Anda Akan diarahkan ke menu dashboard")
		time.Sleep(time.Second * 2)
		app.DasboardUser()

	}

}

func (app *App) UpdateKegiatan() {

}
