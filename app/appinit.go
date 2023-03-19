package app

import (
	"fmt"

	"github.com/ropel12/Latihan/entity"
	"github.com/ropel12/Latihan/repository"
)

type App struct {
	usersRepo    repository.UserInterface
	KegiatanRepo repository.KegiatanInterface
	RencanaRepo repository.RencanaInterface
	Session      map[string]entity.User
}

func NewApp(userRepo repository.UserInterface, KegiatanRepo repository.KegiatanInterface, RencanaRepo repository.RencanaInterface) *App {
	return &App{
		usersRepo:    userRepo,
		KegiatanRepo: KegiatanRepo,
		RencanaRepo: RencanaRepo,
		Session:      make(map[string]entity.User, 0),
	}
}

func (app *App) Start() {
	fmt.Print("\x1bc")
	app.HomePage()
}
