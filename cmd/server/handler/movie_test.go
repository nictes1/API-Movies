package handler

import (
	"api-movies/internal/domain"
	"testing"
	"time"
)

var movie_test = []domain.Movie{
	{
		ID:           1,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "El encargado",
		Rating:       7,
		Awards:       3,
		Release_date: time.Layout,
		Length:       180,
		Genre_id:     2,
	},
	{
		ID:           2,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "Las tortugas Ninjas",
		Rating:       9,
		Awards:       22,
		Release_date: time.Layout,
		Length:       156,
		Genre_id:     3,
	},
	{
		ID:           3,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "Rocky 2",
		Rating:       7,
		Awards:       9,
		Release_date: time.Layout,
		Length:       190,
		Genre_id:     5,
	},
}

func serverMovies()

func Test_GetAll(t *testing.T) {

}
