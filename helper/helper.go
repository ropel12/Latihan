package helper

import (
	"fmt"
	"time"

	"github.com/ropel12/Latihan/entity"
)

func CheckIfError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetUser(user map[string]entity.User) string {
	for _, val := range user {
		return val.Username
	}
	return ""
}

func PrintData(datas interface{}) {

	if datas1, ok := datas.([]entity.Kegiatan); ok {
		var i = 1
		for _, val := range datas1 {
			fmt.Printf("%d. Nama Kegiatan : %s  Tanggal Kegiatan: %s\n", i, val.NamaKegiatan, val.WaktuKegiatan.Format("2006-01-02 15:04:05"))
			i++
		}
	} else if datas2, ok := datas.([]entity.Rencana); ok {
		var i = 1
		for _, val := range datas2 {
			fmt.Printf("%d. Nama Rencana: %s  Tanggal Rencana:%s\n", i, val.NamaRencana, val.WaktuRencana.Format("2006-01-02 15:04:05"))
			i++
		}
	}
}

func ConvertStringToTime(tm string) time.Time {
	tmi, _ := time.Parse("2006-01-02 15:04:05", tm)
	return tmi
}
