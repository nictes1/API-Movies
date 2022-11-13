package movie

import (
	"api-movies/internal/domain"
	"context"
	"errors"
)

type MockMoviesService struct {
	DataMock []domain.Movie
	Error    string
}

func (ms *MockMoviesService) GetAll(ctx context.Context) ([]domain.Movie, error) {
	if ms.Error != "" {
		return []domain.Movie{}, errors.New(ms.Error)
	}
	return ms.DataMock, nil
}

func (ms *MockMoviesService) GetAllMoviesByGenre(ctx context.Context, genreID int) ([]domain.Movie, error) {
	if ms.Error != "" {
		return []domain.Movie{}, errors.New(ms.Error)
	}

	var result = []domain.Movie{}
	for _, movie := range ms.DataMock {
		if movie.Genre_id == genreID {
			result = append(result, movie)
		}
	}

	return result, nil
}

func (ms *MockMoviesService) GetMovieByID(ctx context.Context, id int) (domain.Movie, error) {
	if ms.Error != "" {
		return domain.Movie{}, errors.New(ms.Error)
	}
	for _, movie := range ms.DataMock {
		if id == movie.ID {
			return movie, nil
		}
	}
	return domain.Movie{}, nil
}

func (ms *MockMoviesService) Save(ctx context.Context, movie domain.Movie) (domain.Movie, error) {
	if ms.Error != "" {
		return domain.Movie{}, errors.New(ms.Error)
	}

	for _, m := range ms.DataMock {
		if movie.ID == m.ID {
			return domain.Movie{}, errors.New("already exists ID")
		}
	}

	return movie, nil
}

func (ms *MockMoviesService) Update(ctx context.Context, m domain.Movie, id int) (domain.Movie, error) {
	if ms.Error != "" {
		return domain.Movie{}, errors.New(ms.Error)
	}

	movie, err := ms.GetMovieByID(ctx, id)
	if err != nil {
		return domain.Movie{}, err
	}

	if m.Title != "" {
		movie.Title = m.Title
	}
	if m.Rating != 0 {
		movie.Rating = m.Rating
	}
	if m.Awards != 0 {
		movie.Awards = m.Awards
	}
	if m.Release_date != "" {
		movie.Release_date = m.Release_date
	}
	if m.Length != 0 {
		movie.Length = m.Length
	}
	if m.Genre_id != 0 {
		movie.Genre_id = m.Genre_id
	}

	return movie, nil
}

func (ms *MockMoviesService) Delete(ctx context.Context, id int64) error {
	if ms.Error != "" {
		return errors.New(ms.Error)
	}

	_, err := ms.GetMovieByID(ctx, int(id))
	if err != nil {
		return errors.New("movie to delete doesnt exist")
	}
	for i, m := range ms.DataMock {
		if int(id) == m.ID {
			ms.DataMock = append(ms.DataMock[:i], ms.DataMock[i+1:]...)
		}
	}
	return nil
}
