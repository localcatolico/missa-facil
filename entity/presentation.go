package entity

type Presentation struct {
	Title  string `json:"title"`
	Prayer Prayer `json:"prayer,omitempty"`
	Songs  []Song `json:"songs"`
}

func NewPresentation(title string, prayer Prayer, songs []Song) (*Presentation, error) {
	p := &Presentation{
		Title:  title,
		Prayer: prayer,
		Songs:  songs,
	}

	if err := p.Validade(); err != nil {
		return p, err
	}

	return p, nil
}

func (p *Presentation) Validade() error {
	if p.Title == "" {
		return ErrInvalidEntity
	}
	return nil
}
