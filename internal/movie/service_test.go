package movie

import (
	"api-movies/internal/domain"
	"api-movies/pkg/test/mocks/movie"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CreateMovies(t *testing.T) {
	mockService := movie.MockMoviesRepository{
		DataMock: []domain.Movie{},
		Error:    "",
	}

	movie_test := domain.Movie{
		ID:           1,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "PeliTest",
		Rating:       111,
		Awards:       222,
		Release_date: time.Layout,
		Length:       333,
		Genre_id:     444,
	}

	service := NewService(&mockService)
	result, err := service.Save(context.TODO(), movie_test)
	assert.Nil(t, err)
	assert.Equal(t, movie_test, result)

}

func Test_CreateMoviesFail(t *testing.T) {
	mockService := movie.MockMoviesRepository{
		DataMock: []domain.Movie{},
		Error:    "",
	}

	movie_test := domain.Movie{
		ID:           1,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "PeliTest",
		Rating:       111,
		Awards:       222,
		Release_date: time.Layout,
		Length:       333,
		Genre_id:     444,
	}

	movie_test_fail := domain.Movie{
		ID:           1,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "PeliTest",
		Rating:       111,
		Awards:       222,
		Release_date: time.Layout,
		Length:       333,
		Genre_id:     444,
	}

	service := NewService(&mockService)
	result, err := service.Save(context.TODO(), movie_test)
	assert.Nil(t, err)
	assert.NotEqual(t, movie_test_fail, result)

}
