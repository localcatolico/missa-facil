package usecase_test

import (
	"testing"
)

func TestCreatePresentation(t *testing.T) {
	tests := []struct {
		test     string
		title    string
		prayerId string
		songsId  []string
		err      error
	}{
		{
			test:     "sada",
			title:    "My sides",
			prayerId: "1",
			songsId:  []string{"123", "124"},
			err:      nil,
		},
	}

	// service := usecase.NewUsecase(Seed())
	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			// prayer, _ := service.GetPrayer(tt.prayerId)
			// songs := service.GetSongs()
			// service.CreatePresentation(tt.title, prayer, songs)
		})
	}
}
