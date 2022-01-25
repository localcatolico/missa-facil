package usecase

import (
	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/rafaelbmateus/slides-gospel/google"
	"golang.org/x/oauth2"
)

func (me *Usecase) CreatePresentation(oauth2 *oauth2.Config, token *oauth2.Token,
	title string, prayer entity.Prayer, songs []entity.Song) (*entity.Presentation, error) {

	p := &entity.Presentation{
		Title:  title,
		Prayer: prayer,
		Songs:  songs,
	}

	if _, err := google.CreatePresentation(p, oauth2, token); err != nil {
		return nil, err
	}

	return p, nil
}
