package migration

import (
	"database/sql"

	"github.com/ropel12/Latihan/helper"
)

func Migration(db *sql.DB) error {
	_, err := db.Exec(`
  CREATE TABLE kegiatan_has_rencana (
    kegiatan_id_kegiatan INT NOT NULL,
    kegiatan_id_user INT NOT NULL,
    rencana_id_rencana INT NOT NULL,
    rencana_user_id_user INT NOT NULL,

    FOREIGN KEY (kegiatan_id_kegiatan, kegiatan_id_user)
    REFERENCES kegiatan(id_kegiatan, id_user)

    FOREIGN KEY (rencana_id_rencana,rencana_user_id_user)
    REFERENCES rencana(id_rencana,id_user)
) ENGINE=INNODB;
`)
	return helper.CheckIfError(err)
}
