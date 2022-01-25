package entity

type Songs struct {
	Songs []Song `json:"songs"`
}

type Song struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Artist string  `json:"artist"`
	Slides []Slide `json:"slides,omitempty"`
}

func NewSong(id, name, artist string, slides []Slide) (*Song, error) {
	s := &Song{
		ID:     id,
		Name:   name,
		Artist: artist,
		Slides: slides,
	}

	if err := s.Validade(); err != nil {
		return s, err
	}

	return s, nil
}

func (s *Song) Validade() error {
	if s.ID == "" || s.Name == "" || s.Artist == "" {
		return ErrInvalidEntity
	}

	return nil
}
