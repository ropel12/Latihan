package kegiatantest

import (
	"database/sql"
	"testing"
	"time"

	"github.com/ropel12/Latihan/entity"
	"github.com/ropel12/Latihan/repository"
	"github.com/ropel12/Latihan/repository/Testing"
	"github.com/stretchr/testify/assert"
)

var Db *sql.DB
var RepoKegiatan repository.KegiatanInterface
var newid int64
var idkegiatan int

func TestMain(m *testing.M) {
	Db = Testing.InitDb()
	RepoKegiatan = repository.InitKegiatanRepo(Db)
	res, _ := Db.Exec("INSERT INTO users (username,password) values(?,?)", "satrio", "12345")
	id, _ := res.LastInsertId()
	newid = id
	m.Run()
	// Db.Exec("DELETE from users where username='satrio'")
	// Db.Exec("DELETE from kegiatan where id_user=?", newid)
}
func TestSuccessCreateKegiatan(t *testing.T) {
	var kegiatan = entity.Kegiatan{NamaKegiatan: "Belajar Scala", Userid: int(newid)}
	err := RepoKegiatan.Create(kegiatan)
	assert.Nil(t, err, "Test Case Success")
}
func TestFailedCreateKegiatan(t *testing.T) {
	var user = entity.Kegiatan{NamaKegiatan: "Belajar Scala"}
	err := RepoKegiatan.Create(user)
	assert.Error(t, err, "Test Case Failed")
}

func TestSuccessGetAllKegiatan(t *testing.T) {
	data, _ := RepoKegiatan.GetAll(int(newid))
	assert.NotEqual(t, 0, len(data))
	idkegiatan = data[0].Idkegiatan
}
func TestFailedGetAllKegiatan(t *testing.T) {
	res, _ := RepoKegiatan.GetAll(999)
	assert.Equal(t, 0, len(res), "Test Case Failed Kegiatan")
}
func TestSuccessUpdateKegiatan(t *testing.T) {
	var user = entity.Kegiatan{NamaKegiatan: "Belajar Scalarr", WaktuKegiatan: time.Now()}
	err := RepoKegiatan.UpdateKegiatan(user, idkegiatan)
	assert.Nil(t, err, "Test Case Success Update Kegiatan")
}
func TestFailedUpdateKegiatan(t *testing.T) {
	var user = entity.Kegiatan{NamaKegiatan: "Belajar Scalarr"}
	err := RepoKegiatan.UpdateKegiatan(user, 19992)
	assert.Error(t, err, "Test Case Failed Update Kegiatan")
}
func TestSuccessFindKegiatanByName(t *testing.T) {
	res, _ := RepoKegiatan.FindKegiatanByName("Belajar Scalarr", int(newid))
	assert.Equal(t, "Belajar Scalarr", res.NamaKegiatan)
}
func TestFailedFindKegiatanByName(t *testing.T) {
	res, err := RepoKegiatan.FindKegiatanByName("rrr", int(newid))
	assert.NotEqual(t, "Belajar Scalarr", res.NamaKegiatan)
	assert.Error(t, err)
}

func TestSuccessDeleteKegiatan(t *testing.T) {
	err := RepoKegiatan.DeleteKegiatan(idkegiatan)
	assert.Nil(t, err, "Test Case Success Delete Kegiatan")
}

func TestFailedDeleteKegiatan(t *testing.T) {
	err := RepoKegiatan.DeleteKegiatan(999999999)
	assert.Error(t, err, "Test Case Failed Delete Kegiatan")
}
