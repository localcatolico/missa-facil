package entity

type Prayers struct {
	Prayers []Prayer `json:"prayers"`
}

type Prayer struct {
	ID      string   `json:"id"`
	Prayer  string   `json:"name"`
	Content []string `json:"content"`
}

func NewPrayer(id string, prayer string, content []string) *Prayer {
	return &Prayer{
		ID:      id,
		Prayer:  prayer,
		Content: content,
	}
}
