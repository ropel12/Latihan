package repository

import (
	"database/sql"
	"errors"

	"github.com/ropel12/Latihan/entity"
)

type RencanaRepo struct {
	db *sql.DB
}

func InitRencanaRepo(db *sql.DB) RencanaInterface {
	return &RencanaRepo{db}
}

func (r *RencanaRepo) FindByRencanaId(rencana int) (entity.Rencana, error) {
	var res entity.Rencana
	err := r.db.QueryRow("SELECT id_rencana, nama_rencana, waktu_rencana, updated_at, deleted_at, id_user FROM rencana WHERE id_rencana=?", rencana).Scan(
		&res.IdRencana, &res.NamaRencana, &res.WaktuRencana, &res.UpdateAt, &res.DeleteAt, &res.IdUser,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, errors.New("ID Rencana Tidak Ditemukan")
		}
		return res, err
	}
	return res, nil
}


func (r *RencanaRepo) FindRencanaByUID(IdUser int) ([]entity.Rencana, error) {
	var res []entity.Rencana
	rows, err := r.db.Query("SELECT id_rencana, nama_rencana, waktu_rencana, updated_at, id_user FROM rencana WHERE id_user=?" , IdUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var row entity.Rencana
		
		err := rows.Scan(&row.IdRencana, &row.NamaRencana, &row.WaktuRencana, &row.UpdateAt,  &row.IdUser)
		if err != nil {
			return nil, err
		}
		
		res = append(res,row)
	}
	return res, nil
}
