package usecase

import "github.com/rafaelbmateus/slides-gospel/database"

type Usecase struct {
	Database database.Reader
}

func NewUsecase(db database.Reader) *Usecase {
	return &Usecase{
		Database: db,
	}
}
