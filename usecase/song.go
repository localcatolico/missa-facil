package usecase

import (
	"github.com/rafaelbmateus/slides-gospel/entity"
)

func (me *Usecase) GetSongs() []entity.Song {
	return me.Database.GetSongs()
}

func (me *Usecase) GetSong(songID string) (entity.Song, error) {
	return me.Database.GetSong(songID)
}
