package repository

import "github.com/ropel12/Latihan/entity"

type RencanaInterface interface {
	FindByRencanaId(rencana int) ( entity.Rencana, error)
	FindRencanaByUID(IdUser int) ( []entity.Rencana, error)
	CreateRencana(data entity.Rencana) error
	// UpdateRencana(data entity.Rencana) error
	// DeleteRencana(data entity.Rencana) error
}