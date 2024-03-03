package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/stretchr/testify/assert"
)

func TestPrayer(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		prayer  string
		content []string
		err     error
	}{
		{
			name:    "prayer with success",
			id:      "123",
			prayer:  "prayer name",
			content: []string{"hello"},
			err:     nil,
		},
		{
			name:    "id is required",
			id:      "",
			prayer:  "prayer name",
			content: []string{"hello"},
			err:     entity.ErrInvalidEntity,
		},
		{
			name:    "prayer name is required",
			id:      "123",
			prayer:  "",
			content: []string{"hello"},
			err:     entity.ErrInvalidEntity,
		},
		{
			name:    "with slides empty",
			id:      "123",
			prayer:  "prayer name",
			content: []string{"hello"},
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := entity.NewPrayer(tt.id, tt.prayer, tt.content)
			assert.Equal(t, tt.id, p.ID)
			assert.Equal(t, tt.prayer, p.Prayer)
			assert.Equal(t, tt.content, p.Content)
		})
	}
}
