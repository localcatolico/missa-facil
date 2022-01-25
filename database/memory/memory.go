package memory

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rafaelbmateus/slides-gospel/entity"
)

type Memory struct {
	Prayers entity.Prayers
	Songs   entity.Songs
}

func NewMemory(prayersPath, songsPath string) (*Memory, error) {
	prayers, err := loadPrayers(prayersPath)
	if err != nil {
		return nil, err
	}

	songs, err := loadSongs(songsPath)
	if err != nil {
		return nil, err
	}

	return &Memory{
		Prayers: prayers,
		Songs:   songs,
	}, nil
}

func (me *Memory) GetPrayers() []entity.Prayer {
	return me.Prayers.Prayers
}

func (me *Memory) GetPrayer(id string) (entity.Prayer, error) {
	for _, p := range me.Prayers.Prayers {
		if p.ID == id {
			return p, nil
		}
	}

	return entity.Prayer{}, entity.ErrNotFoundEntity
}

func (me *Memory) GetSongs() []entity.Song {
	return me.Songs.Songs
}

func (me *Memory) GetSong(id string) (entity.Song, error) {
	for _, p := range me.Songs.Songs {
		if p.ID == id {
			return p, nil
		}
	}

	return entity.Song{}, entity.ErrNotFoundEntity
}

func loadPrayers(path string) (entity.Prayers, error) {
	prayers := entity.Prayers{}

	file, err := readFile(path)
	if err != nil {
		return prayers, err
	}

	err = json.Unmarshal([]byte(file), &prayers)
	if err != nil {
		return prayers, err
	}
	return prayers, nil
}

func loadSongs(path string) (entity.Songs, error) {
	songs := entity.Songs{}

	file, err := readFile(path)
	if err != nil {
		return songs, err
	}

	err = json.Unmarshal([]byte(file), &songs)
	if err != nil {
		return songs, err
	}

	return songs, nil
}

func readFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
