package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/stretchr/testify/assert"
)

func TestPresentation(t *testing.T) {
	tests := []struct {
		test   string
		title  string
		prayer entity.Prayer
		songs  []entity.Song
		err    error
	}{
		{
			test:   "with success",
			title:  "presentation title",
			prayer: entity.Prayer{},
			songs:  []entity.Song{},
			err:    nil,
		},
		{
			test:   "without title",
			title:  "",
			prayer: entity.Prayer{},
			songs:  []entity.Song{},
			err:    entity.ErrInvalidEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			p, err := entity.NewPresentation(tt.title, tt.prayer, tt.songs)
			assert.Equal(t, tt.title, p.Title)
			assert.Equal(t, tt.prayer, p.Prayer)
			assert.Equal(t, tt.songs, p.Songs)
			assert.Equal(t, tt.err, err)
		})
	}
}
