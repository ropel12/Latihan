package migration

import (
	"github.com/ropel12/Latihan/db"
	"github.com/ropel12/Latihan/helper"
)

func Migration() error {
	db := db.InitDb()
	_, err := db.Exec(`
  CREATE TABLE IF NOT EXISTS users (
    id_user int NOT NULL AUTO_INCREMENT,
    username varchar(100) NOT NULL,
    password varchar(32) NOT NULL,
    status_account tinyint DEFAULT '1',
    PRIMARY KEY (id_user),
    UNIQUE KEY id_user_UNIQUE (id_user)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

`)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS kegiatan (
  id_kegiatan int NOT NULL AUTO_INCREMENT,
  nama_kegiatan varchar(100) DEFAULT NULL,
  waktu_kegiatan datetime DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  id_user int NOT NULL,
  PRIMARY KEY (id_kegiatan,id_user),
  UNIQUE KEY id_kegiatan_UNIQUE (id_kegiatan),
  KEY fk_kegiatan_user_idx (id_user),
  CONSTRAINT fk_kegiatan_user FOREIGN KEY (id_user) REFERENCES users (id_user),
  CONSTRAINT fk_users_1 FOREIGN KEY (id_user) REFERENCES users (id_user)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;`)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS rencana (
    id_rencana int NOT NULL AUTO_INCREMENT,
    nama_rencana varchar(45) DEFAULT NULL,
    waktu_rencana datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL,
    id_user int NOT NULL,
    PRIMARY KEY (id_rencana,id_user),
    UNIQUE KEY id_rencana_UNIQUE (id_rencana),
    KEY fk_rencana_user1_idx (id_user),
    CONSTRAINT fk_rencana_user1 FOREIGN KEY (id_user) REFERENCES users (id_user)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;`)

	_, err = db.Exec(`CREATE TABLE kegiatan_has_rencana (
    kegiatan_id_kegiatan int NOT NULL,
    kegiatan_id_user int NOT NULL,
    rencana_id_rencana int NOT NULL,
    rencana_user_id_user int NOT NULL,
    KEY fk_kegiatan_has_rencana_rencana1_idx (rencana_id_rencana,rencana_user_id_user),
    KEY fk_kegiatan_has_rencana_kegiatan1_idx (kegiatan_id_kegiatan,kegiatan_id_user),
    CONSTRAINT fk_kegiatan_has_rencana_kegiatan1 FOREIGN KEY (kegiatan_id_kegiatan, kegiatan_id_user) REFERENCES kegiatan (id_kegiatan, id_user),
    CONSTRAINT fk_kegiatan_has_rencana_rencana1 FOREIGN KEY (rencana_id_rencana, rencana_user_id_user) REFERENCES rencana (id_rencana, id_user)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;`)
	defer db.Close()
	return helper.CheckIfError(err)
}
