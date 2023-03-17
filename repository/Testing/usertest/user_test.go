package usertest

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/ropel12/Latihan/entity"
	"github.com/ropel12/Latihan/repository"
	"github.com/ropel12/Latihan/repository/Testing"
	"github.com/stretchr/testify/assert"
)

var Db *sql.DB
var RepoUser repository.UserInterface

func TestMain(m *testing.M) {
	Db = Testing.InitDb()
	RepoUser = repository.InitUserRepo(Db)
	m.Run()
	Db.Exec("delete from users WHERE username=?", "satrio")
}
func TestSuccessCreateUser(t *testing.T) {
	var user = entity.User{Username: "satrio", Password: "12345"}
	err := RepoUser.CreateUser(user)
	assert.Nil(t, err, "Test Case Success")
}
func TestFailedCreateUser(t *testing.T) {
	var user = entity.User{Username: "satrio"}
	err := RepoUser.CreateUser(user)
	err = errors.New("Gagal")
	assert.Error(t, err, "Test Case Failed")
}

func TestSuccessUpdateUser(t *testing.T) {
	var user = entity.User{Username: "satrio", Password: "123452", StatusAkun: 0}
	err := RepoUser.UpdateUser(user)
	assert.Nil(t, err, "Test Case Update User")
}
func TestFailedUpdateUser(t *testing.T) {
	var user = entity.User{Username: "zzzzzzzz", Password: "12345"}
	err := RepoUser.UpdateUser(user)
	assert.Error(t, err, "Test Case Failed Update")
}
func TestSuccessFindByUsername(t *testing.T) {
	datas, _ := RepoUser.FindByUsername("satrio")
	assert.NotNil(t, datas, "Test Case Success FindByusername")
}

func TestFailedFindByUsername(t *testing.T) {
	_, err := RepoUser.FindByUsername("zzzzzz")
	assert.Error(t, err, "Test Case Failed FindByusername")
}
