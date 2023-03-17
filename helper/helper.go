package helper

import (
	"fmt"

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
			fmt.Printf("%d. %s  %s ", i, val.NamaKegiatan, val.WaktuKegiatan)
			i++
		}
	} else if datas2, ok := datas.([]entity.Rencana); ok {
		var i = 1
		for _, val := range datas2 {
			fmt.Printf("%d. %s  %s", i, val.NamaRencana, val.WaktuRencana)
			i++
		}
	}
}
