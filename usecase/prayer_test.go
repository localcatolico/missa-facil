package usecase_test

import (
	"testing"

	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/rafaelbmateus/slides-gospel/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetPrayers(t *testing.T) {
	service := usecase.NewUsecase(Seed())
	songs := service.GetPrayers()
	assert.Equal(t, 3, len(songs))
}

func TestGetPrayer(t *testing.T) {
	tests := []struct {
		test     string
		prayerId string
		err      error
	}{
		{
			test:     "prayer found with success",
			prayerId: "1",
			err:      nil,
		},
		{
			test:     "prayer not found",
			prayerId: "99",
			err:      entity.ErrNotFoundEntity,
		},
	}
	service := usecase.NewUsecase(Seed())
	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			_, err := service.GetPrayer(tt.prayerId)
			assert.Equal(t, tt.err, err)
		})
	}
}
