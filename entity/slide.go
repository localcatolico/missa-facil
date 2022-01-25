package entity

type Slides struct {
	Slides []Slide `json:"slides"`
}

type Slide struct {
	Content string `json:"content"`
}

func NewSlide(content string) (*Slide, error) {
	s := &Slide{
		Content: content,
	}

	if err := s.Validade(); err != nil {
		return s, err
	}

	return s, nil
}

func (s *Slide) Validade() error {
	if s.Content == "" {
		return ErrInvalidEntity
	}

	return nil
}
