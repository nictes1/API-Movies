package domain

type Movie struct {
	ID          int      `json:"id"`
	Title       []string `json:"title"`
	Rating      float32  `json:"rating"`
	Awards      int      `json:"awards"`
	Relese_date string   `json:"relese_date"`
	Lenght      int      `json:"lenght"`
	Genre_id    int      `json:"genre_id"`
}
