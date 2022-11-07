package domain

type Movie struct {
	ID           int     `json:"id"`
	Created_at   string  `json:"create_at"`
	Updated_at   string  `json:"update_at"`
	Title        string  `json:"title"`
	Rating       float32 `json:"rating"`
	Awards       int     `json:"awards"`
	Release_date string  `json:"release_date"`
	Length       *int    `json:"length"`
	Genre_id     *int    `json:"genre_id"`
}
