package movie

import (
	"api-movies/internal/domain"
	"context"
	"errors"
	"fmt"
)

type MockMoviesRepository struct {
	DataMock []domain.Movie
	Error    string
}

func (ms *MockMoviesRepository) GetAll(ctx context.Context) ([]domain.Movie, error) {
	if ms.Error != "" {
		return nil, fmt.Errorf(ms.Error)
	}
	return ms.DataMock, nil
}

func (ms *MockMoviesRepository) GetAllMoviesByGenre(ctx context.Context, id int) ([]domain.Movie, error) {
	if ms.Error != "" {
		return nil, fmt.Errorf(ms.Error)
	}
	return ms.DataMock, nil
}

func (ms *MockMoviesRepository) GetMovieByID(ctx context.Context, id int) (domain.Movie, error) {
	for _, m := range ms.DataMock {
		if id == m.ID {
			return m, nil
		}
	}
	return domain.Movie{}, errors.New("movie not found")
}

func (ms *MockMoviesRepository) Exists(ctx context.Context, id int) bool {
	if ms.Error != "" {
		return false
	}
	for _, m := range ms.DataMock {
		if id == m.ID {
			return true
		}
	}
	return false
}

func (ms *MockMoviesRepository) Save(ctx context.Context, movie domain.Movie) (int, error) {
	if ms.Error != "" {
		return 0, fmt.Errorf(ms.Error)
	}
	id := 1
	movie.ID = id
	ms.DataMock = append(ms.DataMock, movie)
	return id, nil
}

func (ms *MockMoviesRepository) Update(ctx context.Context, movie domain.Movie, id int) error {
	if ms.Error != "" {
		return fmt.Errorf(ms.Error)
	}
	for i, m := range ms.DataMock {
		if movie.ID == m.ID {
			ms.DataMock[i] = movie
			return nil
		}
	}
	return errors.New("movie not found")
}

func (ms *MockMoviesRepository) Delete(ctx context.Context, id int64) error {
	if ms.Error != "" {
		return fmt.Errorf(ms.Error)
	}
	for i, m := range ms.DataMock {
		if m.ID == int(id) {
			ms.DataMock = append(ms.DataMock[:i], ms.DataMock[i+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}
