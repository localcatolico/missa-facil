package usecase_test

import (
	"testing"

	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/rafaelbmateus/slides-gospel/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetSongs(t *testing.T) {
	service := usecase.NewUsecase(Seed())
	songs := service.GetSongs()
	assert.Equal(t, 2, len(songs))
}

func TestGetSong(t *testing.T) {
	tests := []struct {
		test   string
		songId string
		err    error
	}{
		{
			test:   "soung found with success",
			songId: "123",
			err:    nil,
		},
		{
			test:   "soung not found",
			songId: "99",
			err:    entity.ErrNotFoundEntity,
		},
	}
	service := usecase.NewUsecase(Seed())
	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			_, err := service.GetSong(tt.songId)
			assert.Equal(t, tt.err, err)
		})
	}
}
