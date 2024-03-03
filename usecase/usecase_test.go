package usecase_test

import (
	"testing"

	"github.com/rafaelbmateus/slides-gospel/database/memory"
	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/rafaelbmateus/slides-gospel/usecase"
	"github.com/stretchr/testify/assert"
)

func Seed() *memory.Memory {
	return &memory.Memory{
		Songs: entity.Songs{
			[]entity.Song{
				{
					ID: "123", Name: "song name", Artist: "artist name",
				},
				{
					ID: "124", Name: "other song", Artist: "by artist",
				},
			},
		},
		Prayers: entity.Prayers{
			[]entity.Prayer{
				{
					ID: "1", Prayer: "prayer one", Content: []string{"hello"},
				},
				{
					ID: "2", Prayer: "prayer two", Content: []string{"world"},
				},
				{
					ID: "3", Prayer: "prayer three", Content: []string{"content"},
				},
			},
		},
	}
}

func TestNewUsecase(t *testing.T) {
	u := usecase.NewUsecase(&memory.Memory{})
	assert.NotNil(t, u)
}
