package memory

import (
	"encoding/json"
	"log"
	"net/http"

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
	res, err := http.Get(path)
	if err != nil {
		log.Fatal("Erro ao fazer a solicitação HTTP:", err)

		return entity.Prayers{}, err
	}

	defer res.Body.Close()

	var prayers entity.Prayers
	if err := json.NewDecoder(res.Body).Decode(&prayers); err != nil {
		log.Fatal("Erro ao decodificar a resposta JSON:", err)

		return entity.Prayers{}, err
	}

	return prayers, nil
}

func loadSongs(path string) (entity.Songs, error) {
	res, err := http.Get(path)
	if err != nil {
		log.Fatal("Erro ao fazer a solicitação HTTP:", err)

		return entity.Songs{}, err
	}

	defer res.Body.Close()

	var songs entity.Songs
	if err := json.NewDecoder(res.Body).Decode(&songs); err != nil {
		log.Fatal("Erro ao decodificar a resposta JSON:", err)

		return entity.Songs{}, err
	}

	return songs, nil
}
