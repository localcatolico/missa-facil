package database

import "github.com/rafaelbmateus/slides-gospel/entity"

type Reader interface {
	GetPrayers() []entity.Prayer
	GetPrayer(id string) (entity.Prayer, error)
	GetSongs() []entity.Song
	GetSong(id string) (entity.Song, error)
}
