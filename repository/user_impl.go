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

func (u *UserRepo) FindByUsername(username string) (entity.User, error) {
	res := entity.User{}
	row, err := u.db.Query("SELECT id_user,username,password,status_account FROM users WHERE username=?", username)
	if err != nil {
		return res, err
	}
	defer row.Close()
	if row.Next() {
		row.Scan(&res.Userid, &res.Username, &res.Password, &res.StatusAkun)
		return res, nil
	}
	return res, errors.New("User Tidak Ditemukan")

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

func (u *UserRepo) UpdateUser(data entity.User, oldusername string) error {
	row, err := u.db.Exec("UPDATE users set username=?, password=?,status_account=? WHERE username=?", data.Username, data.Password, data.StatusAkun, oldusername)
	if err != nil {
		return err
	}

	rowaff, _ := row.RowsAffected()
	if rowaff > 0 {
		return nil
	}
	return errors.New("Data Tidak Berhasil Di Update")
}
