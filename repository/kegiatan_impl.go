package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ropel12/Latihan/entity"
)

type KegiatanRepo struct {
	db *sql.DB
}

func InitKegiatanRepo(db *sql.DB) KegiatanInterface {
	return &KegiatanRepo{db}
}

func (k *KegiatanRepo) GetAll(id int) ([]entity.Kegiatan, error) {
	rows, err := k.db.Query("SELECT id_kegiatan,nama_kegiatan,waktu_kegiatan,updated_at,id_user from kegiatan WHERE id_user=?", id)
	res := make([]entity.Kegiatan, 0)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := entity.Kegiatan{}
		err := rows.Scan(&row.Idkegiatan, &row.NamaKegiatan, &row.WaktuKegiatan, &row.UpdateAt, &row.Userid)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

func (k *KegiatanRepo) Create(data entity.Kegiatan) error {
	row, err := k.db.Exec("INSERT INTO kegiatan (nama_kegiatan,id_user)values(?,?)", data.NamaKegiatan, data.Userid)
	if err != nil {
		return err
	}
	aff, _ := row.RowsAffected()
	if aff > 0 {
		return nil
	}
	return errors.New("Gagal Membuat Kegiatan")
}

func (k *KegiatanRepo) FindKegiatanByName(name string) (res entity.Kegiatan, err error) {
	row, err := k.db.Query("SELECT id_kegiatan,nama_kegiatan,waktu_kegiatan,update_at,id_user from kegiatan WHERE deleted_at = NULL")
	if err != nil {
		return entity.Kegiatan{}, err
	}
	if row.Next() {
		row.Scan(&res.Idkegiatan, &res.NamaKegiatan, &res.WaktuKegiatan, &res.UpdateAt)
		err = nil
		return
	}
	return entity.Kegiatan{}, errors.New("Data Tidak Ditemukan")
}

func (k *KegiatanRepo) UpdateKegiatan(data entity.Kegiatan, id int) error {
	row, err := k.db.Exec("UPDATE kegiatan set nama_kegiatan=?,waktu_kegiatan=?,updated_at=? WHERE id_kegiatan=?", data.NamaKegiatan, data.WaktuKegiatan, time.Now(), id)
	if err != nil {
		return err
	}
	aff, _ := row.RowsAffected()
	if aff > 0 {
		return nil
	}
	return errors.New("Gagal Update Kegiatan")
}

func (k *KegiatanRepo) DeleteKegiatan(id int) error {
	row, err := k.db.Exec("UPDATE kegiatan set deleted_at=? where id_kegiatan=?", time.Now(), id)
	if err != nil {
		return err
	}
	aff, _ := row.RowsAffected()
	if aff > 0 {
		return nil
	}
	return errors.New("Gagal Delete Kegiatan")
}
