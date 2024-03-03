package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/stretchr/testify/assert"
)

func TestSong(t *testing.T) {
	tests := []struct {
		test    string
		id      string
		name    string
		artist  string
		content []string

		err error
	}{
		{
			test:    "with success",
			id:      "123",
			name:    "song name",
			artist:  "artist name",
			content: []string{},
			err:     nil,
		},
		{
			test:    "id is required",
			id:      "",
			name:    "song name",
			artist:  "artist name",
			content: []string{},
			err:     entity.ErrInvalidEntity,
		},
		{
			test:    "name is required",
			id:      "123",
			name:    "",
			artist:  "artist name",
			content: []string{},
			err:     entity.ErrInvalidEntity,
		},
		{
			test:    "artist is required",
			id:      "123",
			name:    "song name",
			artist:  "",
			content: []string{},
			err:     entity.ErrInvalidEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			p := entity.NewSong(tt.id, tt.name, tt.artist, tt.content)
			assert.Equal(t, tt.id, p.ID)
			assert.Equal(t, tt.name, p.Name)
			assert.Equal(t, tt.artist, p.Artist)
			assert.Equal(t, tt.content, p.Content)
		})
	}
}
