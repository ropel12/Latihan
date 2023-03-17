package repository

import "github.com/ropel12/Latihan/entity"

type KegiatanInterface interface {
	GetAll(id int) ([]entity.Kegiatan, error)
	Create(data entity.Kegiatan) error
	FindKegiatanByName(name string) (res entity.Kegiatan, err error)
	UpdateKegiatan(data entity.Kegiatan, id int) error
	DeleteKegiatan(id int) error
}
