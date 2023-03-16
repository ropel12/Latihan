package repository

import (
	"database/sql"
	"errors"

	"github.com/ropel12/Latihan/entity"
)

type UserRepo struct {
	db *sql.DB
}

func InitUserRepo(db *sql.DB) UserInterface {
	return &UserRepo{db}
}

func (u *UserRepo) FindByUsername(username string) (res *entity.User, err error) {
	row, err := u.db.Query("SELECT id_user,username,password,status_account FROM users WHERE username=?", username)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		row.Scan(&res.Userid, &res.Username, &res.Password, &res.StatusAkun)
		err = nil
		return
	}
	err = errors.New("User Tidak Ditemukan")
	res = nil
	return

}

func (u *UserRepo) CreateUser(data entity.User) error {

	row, err := u.db.Exec("INSERT INTO users (username,password) values (?,?)", data.Username, data.Password)
	if err != nil {
		return err
	}
	rowaff, _ := row.RowsAffected()
	if rowaff > 0 {
		return nil
	}
	return errors.New("Gagal Menambahkan User")

}
