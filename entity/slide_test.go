package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/stretchr/testify/assert"
)

func TestSlide(t *testing.T) {
	tests := []struct {
		name    string
		content string
		err     error
	}{
		{
			name:    "with success",
			content: "hello",
			err:     nil,
		},
		{
			name:    "without content",
			content: "",
			err:     entity.ErrInvalidEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := entity.NewSlide(tt.content)
			assert.Equal(t, tt.content, p.Content)
			assert.Equal(t, tt.err, err)
		})
	}
}
