package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/stretchr/testify/assert"
)

func TestPrayer(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		prayer string
		slides []entity.Slide
		err    error
	}{
		{
			name:   "prayer with success",
			id:     "123",
			prayer: "prayer name",
			slides: []entity.Slide{{Content: "hello"}},
			err:    nil,
		},
		{
			name:   "id is required",
			id:     "",
			prayer: "prayer name",
			slides: []entity.Slide{{Content: "hello"}},
			err:    entity.ErrInvalidEntity,
		},
		{
			name:   "prayer name is required",
			id:     "123",
			prayer: "",
			slides: []entity.Slide{{Content: "hello"}},
			err:    entity.ErrInvalidEntity,
		},
		{
			name:   "with slides empty",
			id:     "123",
			prayer: "prayer name",
			slides: []entity.Slide{},
			err:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := entity.NewPrayer(tt.id, tt.prayer, tt.slides)
			assert.Equal(t, tt.id, p.ID)
			assert.Equal(t, tt.prayer, p.Prayer)
			assert.Equal(t, tt.slides, p.Slides)
			assert.Equal(t, tt.err, err)
		})
	}
}
