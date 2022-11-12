package movie

import (
	"api-movies/internal/domain"
	"api-movies/pkg/utils"
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	ErrForzado = errors.New("Error forzado")
)

/*
func TestGetOneWithContext(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movie_test := domain.Movie{
		ID:           1,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "Cars 1",
		Rating:       4,
		Awards:       2,
		Release_date: time.Layout,
		Length:       0,
		Genre_id:     0,
	}

	columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id)
	mock.ExpectQuery(regexp.QuoteMeta(GET_MOVIE)).WithArgs(movie_test.ID).WillReturnRows(rows)

	repo := NewRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	movieResult, err := repo.GetMovieWithContext(ctx, movie_test.ID)
	assert.NoError(t, err)
	assert.Equal(t, movie_test.Title, movieResult.Title)
	assert.Equal(t, movie_test.ID, movieResult.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
*/

func TestExistMovieOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movie_test := domain.Movie{
		ID:           1,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "Cars 1",
		Rating:       4,
		Awards:       2,
		Release_date: time.Layout,
		Length:       0,
		Genre_id:     0,
	}

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(movie_test.ID)
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_MOVIE)).WithArgs(1).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.Exists(context.TODO(), 1)

	assert.True(t, resp)
}

func TestExistMovieFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	mock.ExpectQuery(regexp.QuoteMeta(EXIST_MOVIE)).WithArgs(2).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.Exists(context.TODO(), 2)

	assert.False(t, resp)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movie_test := domain.Movie{
		ID:           1,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "Cars 1",
		Rating:       4,
		Awards:       2,
		Release_date: time.Layout,
		Length:       0,
		Genre_id:     0,
	}

	t.Run("Store Ok", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_MOVIE)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id)
		mock.ExpectQuery(regexp.QuoteMeta(GET_MOVIE)).WithArgs(1).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()

		newID, err := repository.Save(ctx, movie_test)
		assert.NoError(t, err)

		movieResult, err := repository.GetMovieByID(ctx, int(newID))
		assert.NoError(t, err)

		assert.NotNil(t, movieResult)
		assert.Equal(t, movie_test.ID, movieResult.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Store Fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_MOVIE)).WillReturnError(ErrForzado)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Save(ctx, movie_test)

		assert.EqualError(t, err, ErrForzado.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func Test_RepositoryGetAllOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Columns
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	rows := sqlmock.NewRows(columns)
	movies := []domain.Movie{{ID: 1, Title: "Avatar", Rating: 22, Awards: 99, Length: 0, Genre_id: 1}, {ID: 2, Title: "Simpson", Rating: 33, Awards: 11, Length: 2, Genre_id: 2}}

	for _, movie := range movies {
		rows.AddRow(movie.ID, movie.Title, movie.Rating, movie.Awards, movie.Length, movie.Genre_id)
	}

	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_MOVIES)).WillReturnRows(rows)

	repo := NewRepository(db)
	resultMovies, err := repo.GetAll(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, movies, resultMovies)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_RepositoryGetAllFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// Columns
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	rows := sqlmock.NewRows(columns)
	movies := []domain.Movie{{ID: 1, Title: "Avatar", Rating: 22, Awards: 99, Length: 0, Genre_id: 1}, {ID: 2, Title: "Simpson", Rating: 33, Awards: 11, Length: 2, Genre_id: 2}}

	for _, movie := range movies {
		rows.AddRow(movie.ID, movie.Title, movie.Rating, movie.Awards, movie.Length, movie.Genre_id)
	}

	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_MOVIES)).WillReturnError(ErrForzado)

	repo := NewRepository(db)
	resultMovies, err := repo.GetAll(context.TODO())

	assert.EqualError(t, err, ErrForzado.Error())
	assert.Empty(t, resultMovies)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_UpdateOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movie_test := domain.Movie{
		ID:           1,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
		Title:        "Cars 1",
		Rating:       4,
		Awards:       2,
		Release_date: time.Layout,
		Length:       0,
		Genre_id:     0,
	}

	m := movie_test

	mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_MOVIE)).
		ExpectExec().WithArgs(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id, m.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewRepository(db)
	err = repo.Update(context.TODO(), m, 1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_DeleteOK(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

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

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_MOVIE))
	mock.ExpectExec(regexp.QuoteMeta(DELETE_MOVIE)).WithArgs(movie_test.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewRepository(db)
	// Act
	err = repo.Delete(context.TODO(), 1)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_DeleteFail(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

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

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_MOVIE)).ExpectExec().WithArgs(movie_test.ID).WillReturnError(ErrForzado)

	repo := NewRepository(db)

	// act
	err = repo.Delete(context.TODO(), int64(movie_test.ID))

	// assert
	assert.EqualError(t, err, ErrForzado.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_DeleteFailRowsAffected(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

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

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_MOVIE)).ExpectExec().WithArgs(movie_test.ID).WillReturnResult(sqlmock.NewResult(1, 2))

	repo := NewRepository(db)

	// act
	err = repo.Delete(context.TODO(), int64(movie_test.ID))

	// assert
	assert.EqualError(t, err, "error: no affected rows")
}

////TEXDB 	//////////////////////////////////////////////////////////////////
func Test_RepositorySave_txdb(t *testing.T) {
	db := utils.InitTxDB()
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.TODO()
	// (&m.Title, &m.Rating, &m.Awards, &m.Length, &m.Genre_id

	movieExp := domain.Movie{
		Title:        "Título ficticio",
		Rating:       2,
		Awards:       3,
		Length:       3,
		Genre_id:     2,
		Release_date: "2022-11-09 00:00:00",
	}

	// Act
	id, err := repo.Save(ctx, movieExp)
	assert.NoError(t, err)

	movieExp.ID = int(id)
	movies, err := repo.GetMovieByID(context.TODO(), int(id))

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, movies)
	assert.NoError(t, err)
	assert.Equal(t, movieExp.ID, movies.ID)
}
