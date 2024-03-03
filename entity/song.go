package entity

type Songs struct {
	Songs []Song `json:"musics"`
}

type Song struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Artist  string   `json:"artist"`
	Content []string `json:"content"`
}

func NewSong(id, name, artist string, content []string) *Song {
	return &Song{
		ID:      id,
		Name:    name,
		Artist:  artist,
		Content: content,
	}
}
