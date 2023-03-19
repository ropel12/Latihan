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
	fmt.Print("\x1bc")
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
	fmt.Println("8.Hapus Akun")
	fmt.Println("Masukan Pilihan Anda:")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		app.ListKegiatan()
	case 2:
		app.ListRencana()
	case 3:
		app.FormTambahKegiatan()
	case 4:
		app.FormTambahRencana()
	case 5:
		app.UpdateProfile()
	case 6:
	case 7:
		app.Logout()
	case 8:
		app.DeleteAccount()
	}
}

func (app *App) UpdateProfile() {
	var choice int
	var oldusername string
	defer func() {
		fmt.Print("Jika Ingin Kembali Tekan 9 : ")
		fmt.Scanln(&choice)
		if choice == 9 {
			app.DasboardUser()
		}
		app.UpdateProfile()
	}()
	fmt.Print("\x1bc")
	key := helper.GetUser(app.Session)
	oldusername = app.Session[key].Username
	var username, password string
	fmt.Printf("Username  Lama : %s\n", app.Session[key].Username)
	fmt.Printf("Password Lama : %s\n\n", app.Session[key].Password)
	fmt.Print("Masukan Username Baru: ")
	fmt.Scanln(&username)
	_, err := app.usersRepo.FindByUsername(username)
	if err == nil {
		fmt.Println("Username Telah Terdaftar Silahkan Gunakan Yang Lain")
		time.Sleep(time.Second * 2)
		app.UpdateProfile()
		return
	}
	fmt.Print("Masukan Password Baru: ")
	fmt.Scanln(&password)
	err1 := app.usersRepo.UpdateUser(entity.User{Username: username, Password: password, StatusAkun: 1}, app.Session[key].Username)
	if err1 != nil {
		fmt.Println(err1.Error())
		fmt.Println("Anda akan diarahkan ke halaman dashboard")
		time.Sleep(time.Second * 3)
		app.DasboardUser()
		return
	}
	fmt.Println("Berhasil Update Profile")
	fmt.Println("Anda akan diarahkan ke halaman dashboard")
	newdata, _ := app.usersRepo.FindByUsername(username)
	delete(app.Session, oldusername)
	fmt.Println(oldusername)
	app.Session[username] = newdata
	time.Sleep(time.Second * 3)
	app.DasboardUser()
	return

}

func (app *App) FormTambahKegiatan() {
	fmt.Print("\x1bc")
	key := helper.GetUser(app.Session)
	var namakegiatan, choice string
	fmt.Println("=================FORM TAMBAH KEGIATAN===========================")
	fmt.Println()
	fmt.Print("Masukan Nama Kegiatan: ")
	app.Scanner.Scan()
	namakegiatan = app.Scanner.Text()
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
		fmt.Scanln(&choice)
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
	fmt.Println("\x1bc")
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
		fmt.Scanln(&choice)
		if choice == 1 {
			app.FormTambahKegiatan()
			return
		}
		app.DasboardUser()
		return

	}
	helper.PrintData(datas)
	fmt.Println("Jika anda ingin menambahkan lagi masukan angka 1\njika ingin menghapus masukan angka 2\njika ingin mengupdate masukan angka 3")
	fmt.Println("Jika anda ingin kemmbali ke menu dashboard masukan angka 4")
	fmt.Print("Masukan Pilihan : ")
	fmt.Scanln(&choice)
	if choice == 1 {
		app.FormTambahKegiatan()
		return
	} else if choice == 2 {
		app.HapusKegiatan()
		return
	} else if choice == 3 {
		app.UpdateKegiatan()
		return
	} else if choice == 4 {
		app.DasboardUser()
		return
	}
	fmt.Println("Anda akan diarahkan ke halaman dashboard")
	time.Sleep(time.Second * 1)
	app.DasboardUser()
}

func (app *App) HapusKegiatan() {
	fmt.Println("\x1bc")
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
	if len(datas) == 0 {
		app.ListKegiatan()
		return
	}
	if len(datas) != 0 {
		fmt.Println("Silahkan pilih daftar kegiatan yang ingin dihapus")
		helper.PrintData(datas)
		fmt.Print("Masukan data yang ingin dihapus jika ingin banyak tambahkan koma contoh(1,2,3,4,5)")
		fmt.Scanln(&choices)
		fmt.Println()
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
				index += i + 1
			}
			if index < 1 {
				fmt.Println("Masukan Data Yang benar")
				time.Sleep(2 * time.Second)
				app.HapusKegiatan()
				return
			}
			fmt.Println("Berhasil Mengapus Data")
			fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
			fmt.Scanln(&choices)
			if choices == "y" {
				app.HapusKegiatan()
				return
			}
			fmt.Println("Anda Akan diarahkan ke menu dashboard")
			time.Sleep(time.Second * 2)
			app.DasboardUser()
			return
		}
		toint, _ := strconv.Atoi(choices)
		err := app.KegiatanRepo.DeleteKegiatan(datas[toint-1].Idkegiatan)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Masukan Data Yang Benar")
			app.HapusKegiatan()
			return
		}
		fmt.Println("Berhasil Mengapus Data")
		fmt.Print("Apakah Anda Ingin Melanjutkan (y/t)")
		fmt.Scanln(&choices)
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
	key := helper.GetUser(app.Session)
	var choice int
	var choice2 string
	var kegiatan string
	var waktu string
	datas, err := app.KegiatanRepo.GetAll(app.Session[key].Userid)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(time.Second * 2)
		app.DasboardUser()
	}
	fmt.Print("\x1bc")
	fmt.Println("==================================Daftar Kegiatan==========================================")
	helper.PrintData(datas)
	fmt.Println("Silahkan Pilih Daftar Kegiatan Yang Ingin Diubah")
	fmt.Print("Pilih : ")
	fmt.Scanln(&choice)
	fmt.Println("Masukan Nama Kegiatan Baru: ")
	app.Scanner.Scan()
	kegiatan = app.Scanner.Text()
	fmt.Println("Masukan Tanggal Kegiatan Baru contoh( 2023-03-18 00:20:10 ): ")
	app.Scanner.Scan()
	waktu = app.Scanner.Text()
	fmt.Println(helper.ConvertStringToTime(waktu))
	err1 := app.KegiatanRepo.UpdateKegiatan(entity.Kegiatan{NamaKegiatan: kegiatan, WaktuKegiatan: helper.ConvertStringToTime(waktu)}, datas[choice-1].Idkegiatan)
	if err1 != nil {
		fmt.Println(err1.Error())
		fmt.Print("Apakah Ingin Mencoba Ulang ? (y/t): ")
		fmt.Scanln(&choice2)
		if choice2 == "y" {
			app.UpdateKegiatan()
			return
		}
		fmt.Println("Anda Akan Diarahkan Kehalaman Dashboard")
		time.Sleep(3 * time.Second)
		app.DasboardUser()
		return
	}
	fmt.Println("Berhasil Mengupdate Data")
	fmt.Print("Apakah Anda Ingin Update Lagi ? (y/t): ")
	fmt.Scanln(&choice2)
	if choice2 == "y" {
		app.UpdateKegiatan()
		return
	}
	fmt.Println("Anda Akan Diarahkan Kehalaman Dashboard")
	time.Sleep(3 * time.Second)
	app.DasboardUser()
	return

}

func (app *App) DeleteAccount() {
	key := helper.GetUser(app.Session)
	app.usersRepo.UpdateUser(entity.User{Username: app.Session[key].Username, Password: app.Session[key].Password, StatusAkun: 0}, app.Session[key].Username)
	fmt.Println("Berhasil Menghapus Akun Tunggu 3 detik dan akan redirect halaman home")
	time.Sleep(time.Second * 3)

}

func (app *App) ListRencana() {
	key := helper.GetUser(app.Session)
	var choice int
	datas, err := app.RencanaRepo.FindRencanaByUID(app.Session[key].Userid)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(time.Second * 2)
		app.DasboardUser()
	}
	fmt.Printf("Berikut ini daftar rencana dari %s yang telah dibuat\n", app.Session[key].Username)
	if len(datas) == 0 {
		fmt.Print("Anda belum memiliki daftar rencana. Jika Anda ingin menambahkan rencana masukan angka 1: ")
		fmt.Scanln(&choice)
		if choice == 1 {
			app.FormTambahRencana()
		}
		app.DasboardUser()
	}
	helper.PrintData(datas)
	fmt.Print("Jika anda ingin menambahkan lagi masukan angka 1 dan jika ingin menghapus masukan angka 2 dan jika ingin mengupdate masukan angka 3 : ")
	fmt.Scanln(&choice)
	if choice == 1 {
		app.FormTambahRencana()
	} else if choice == 2 {
		// app.HapusRencana()
	} else if choice == 3 {
		// app.UpdateRencana()
	}
	fmt.Println("Anda akan diarahkan ke halaman dashboard")
	time.Sleep(time.Second * 1)
	app.DasboardUser()
}

func (app *App) FormTambahRencana() {
	key := helper.GetUser(app.Session)
	var namarencana, choice string
	var userId int
	fmt.Println("=================FORM TAMBAH RENCANA===========================")
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukan Nama Kegiatan: ")
	fmt.Scan(&namarencana)
	_, err := app.RencanaRepo.FindRencanaByUID(userId)
	if err != nil {
		err2 := app.RencanaRepo.CreateRencana(entity.Rencana{NamaRencana: namarencana, IdRencana: app.Session[key].Userid})
		if err2 != nil {
			fmt.Println(err2.Error())
			app.FormTambahRencana()
		}
		fmt.Println("Tambah Rencana Berhasil Ditambahkan")
		fmt.Print("Ingin Menambahkan data lagi? (y/t)")
		fmt.Scan(&choice)
		if choice == "y" {
			app.FormTambahRencana()
		}
		fmt.Println("Anda akan redirect ke menu utama dalam 3 detik")
		time.Sleep(time.Second * 3)
		app.DasboardUser()
	}
	fmt.Println("Nama Rencana Sudah Ada")
	time.Sleep(time.Second * 2)
	app.FormTambahRencana()

}