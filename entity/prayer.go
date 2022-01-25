package entity

type Prayers struct {
	Prayers []Prayer `json:"prayers"`
}

type Prayer struct {
	ID     string  `json:"id"`
	Prayer string  `json:"name"`
	Slides []Slide `json:"slides"`
}

func NewPrayer(id string, prayer string, slides []Slide) (*Prayer, error) {
	p := &Prayer{
		ID:     id,
		Prayer: prayer,
		Slides: slides,
	}

	if err := p.Validade(); err != nil {
		return p, err
	}

	return p, nil
}

func (p *Prayer) Validade() error {
	if p.ID == "" || p.Prayer == "" {
		return ErrInvalidEntity
	}

	return nil
}
