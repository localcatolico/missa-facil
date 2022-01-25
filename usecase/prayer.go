package usecase

import (
	"github.com/rafaelbmateus/slides-gospel/entity"
)

func (me *Usecase) GetPrayers() []entity.Prayer {
	return me.Database.GetPrayers()
}

func (me *Usecase) GetPrayer(prayerID string) (entity.Prayer, error) {
	return me.Database.GetPrayer(prayerID)
}
