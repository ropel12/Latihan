package repository

import (
	"database/sql"

	"github.com/ropel12/Latihan/entity"
)

type KegiatanRepo struct {
	db *sql.DB
}

func InitKegiatanRepo(db *sql.DB) Kegiatan {
	return &KegiatanRepo{db}
}

func (k *KegiatanRepo) GetAll(id int) ([]entity.Kegiatan, error) {
	rows, err := k.db.Query("SELECT id_kegiatan,nama_kegiatan,waktu_kegiatan,update_at,delete_at,id_user from kegiatan WHERE id_user=?", id)
	res := make([]entity.Kegiatan, 0)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := entity.Kegiatan{}
		err := rows.Scan(&row.Idkegiatan, &row.NamaKegiatan, &row.WaktuKegiatan, &row.UpdateAt, &row.DeleteAt, &row.Userid)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}
