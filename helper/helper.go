package helper

import (
	"fmt"
	"strings"
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
			fmt.Printf("%d. %s  %s\n", i, val.NamaKegiatan, val.WaktuKegiatan)
			i++
		}
	} else if datas2, ok := datas.([]entity.Rencana); ok {
		var i = 1
		for _, val := range datas2 {
			fmt.Printf("%d. %s  %s\n", i, val.NamaRencana, val.WaktuRencana)
			i++
		}
	}
}
func ConvertTimezone(tm string) string {

	zone, _ := time.LoadLocation("Asia/Jakarta")
	t, _ := time.ParseInLocation(tm, "05-05-2019 05:11", zone)
	return t.String()
}

func ConvertStringToTime(tm string) time.Time {

	tmi, _ := time.Parse(strings.Replace(tm, " ", "", -1), tm)
	return tmi
}
