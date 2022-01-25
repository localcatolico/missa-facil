package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/stretchr/testify/assert"
)

func TestSong(t *testing.T) {
	tests := []struct {
		test   string
		id     string
		name   string
		artist string
		slides []entity.Slide

		err error
	}{
		{
			test:   "with success",
			id:     "123",
			name:   "song name",
			artist: "artist name",
			slides: []entity.Slide{},
			err:    nil,
		},
		{
			test:   "id is required",
			id:     "",
			name:   "song name",
			artist: "artist name",
			slides: []entity.Slide{},
			err:    entity.ErrInvalidEntity,
		},
		{
			test:   "name is required",
			id:     "123",
			name:   "",
			artist: "artist name",
			slides: []entity.Slide{},
			err:    entity.ErrInvalidEntity,
		},
		{
			test:   "artist is required",
			id:     "123",
			name:   "song name",
			artist: "",
			slides: []entity.Slide{},
			err:    entity.ErrInvalidEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			p, err := entity.NewSong(tt.id, tt.name, tt.artist, tt.slides)
			assert.Equal(t, tt.id, p.ID)
			assert.Equal(t, tt.name, p.Name)
			assert.Equal(t, tt.artist, p.Artist)
			assert.Equal(t, tt.slides, p.Slides)
			assert.Equal(t, tt.err, err)
		})
	}
}
