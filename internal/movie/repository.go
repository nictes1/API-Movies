package movie

import (
	"api-movies/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Movie, error)
	Get(ctx context.Context, id int) (domain.Movie, error)
	Exists(ctx context.Context, id int) bool
	Save(ctx context.Context, m domain.Movie) (int64, error)
	Update(ctx context.Context, b domain.Movie) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	SAVE_MOVIE     = "INSERT INTO movies (created_at, title, rating, awards, length, genre_id) VALUES (?,?,?,?,?,?);"
	GET_ALL_MOVIES = "SELECT m.id ,m.title, m.rating, m.awards, m.length, m.genre_id FROM movies m;"
	GET_MOVIE      = "SELECT m.id, m.title, m.rating, m.awards, m.length, m.genre_id FROM movies m WHERE m.id=?;"
	UPDATE_MOVIE   = "UPDATE movies SET updated_at=?, title=?, rating=?, awards=?, relese_date=?, length=?, genre_id=? WHERE id=?;"
	DELETE_MOVIE   = "DELETE FROM movies WHERE id=?;"
	EXIST_MOVIE    = "SELECT m.id FROM movies m WHERE m.id=?"
)

func (r *repository) GetAll(ctx context.Context) ([]domain.Movie, error) {
	rows, err := r.db.Query(GET_ALL_MOVIES)
	if err != nil {
		return nil, err
	}

	var movies []domain.Movie

	for rows.Next() {
		var movie domain.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id); err != nil {
			return []domain.Movie{}, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *repository) Get(ctx context.Context, id int) (movie domain.Movie, err error) {

	rows, err := r.db.Query(GET_MOVIE)
	if err != nil {
		return domain.Movie{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&movie.ID, &movie.Created_at, &movie.Updated_at, &movie.Title, &movie.Rating, &movie.Awards, &movie.Release_date, &movie.Length, &movie.Genre_id); err != nil {
			return domain.Movie{}, err
		}

	}
	return movie, nil
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(EXIST_MOVIE, id)

	err := rows.Scan()
	fmt.Println(err)
	return false
}

func (r *repository) Save(ctx context.Context, m domain.Movie) (int64, error) {
	stm, err := r.db.Prepare(SAVE_MOVIE)
	if err != nil {
		return 0, err
	}

	res, err := stm.Exec(time.Now(), &m.Title, &m.Rating, &m.Awards, &m.Length, &m.Genre_id)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Update(ctx context.Context, m domain.Movie) error {
	stm, err := r.db.Prepare(UPDATE_MOVIE)
	if err != nil {
		return err
	}
	res, err := stm.Exec(time.Now(), &m.Title, &m.Rating, &m.Awards, &m.Release_date, &m.Length, &m.Genre_id, &m.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stm, err := r.db.Prepare(DELETE_MOVIE)
	if err != nil {
		return err
	}
	res, err := stm.Exec(id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil || affected == 0 {
		return err
	}
	return nil
}
