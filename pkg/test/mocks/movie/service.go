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

func (m *MockMoviesService) GetAll(ctx context.Context) ([]domain.Movie, error) {
	if m.Error != "" {
		return []domain.Movie{}, errors.New(m.Error)
	}
	return m.DataMock, nil
}

func (m *MockMoviesService) GetAllMoviesByGenre(ctx context.Context, genreID int) ([]domain.Movie, error) {
	if m.Error != "" {
		return []domain.Movie{}, errors.New(m.Error)
	}
	result, err := m.GetAllMoviesByGenre(ctx, genreID)
	if err != nil {
		return []domain.Movie{}, errors.New("no movies of the genre will be found")
	}
	m.DataMock = result
	return m.DataMock, nil
}

func (m *MockMoviesService) GetMovieByID(ctx context.Context, id int) (domain.Movie, error) {
	if m.Error != "" {
		return domain.Movie{}, errors.New(m.Error)
	}
	for _, m := range m.DataMock {
		if id == m.ID {
			return m, nil
		}
	}
	return domain.Movie{}, nil
}

func (ms *MockMoviesService) Create(ctx context.Context, movie domain.Movie) (domain.Movie, error) {
	for _, m := range ms.DataMock {
		if movie.ID == m.ID {
			return domain.Movie{}, errors.New("already exists ID")
		}
	}

	return movie, nil
}

func (ms *MockMoviesService) Update(ctx context.Context, m domain.Movie, id int) (domain.Movie, error) {

	movie, err := ms.GetMovieByID(ctx, id)
	if err != nil {
		return domain.Movie{}, err
	}

	movie.ID = m.ID
	movie.Created_at = m.Created_at
	movie.Updated_at = m.Updated_at
	movie.Title = m.Title
	movie.Rating = m.Rating
	movie.Awards = m.Awards
	movie.Release_date = m.Release_date
	movie.Length = m.Length
	movie.Genre_id = m.Genre_id

	return movie, nil
}

func (ms *MockMoviesService) Delete(ctx context.Context, id int64) error {
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
