package movie

import (
	"api-movies/internal/domain"
	"context"
	"database/sql"
	"errors"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Movie, error)
	GetAllMoviesByGenre(ctx context.Context, genreID int) ([]domain.Movie, error)
	GetMovieByID(ctx context.Context, id int) (domain.Movie, error)
	GetMovieWithContext(ctx context.Context, id int) (domain.Movie, error)
	Save(ctx context.Context, b domain.Movie) (int64, error)
	Exists(ctx context.Context, id int) bool
	Update(ctx context.Context, b domain.Movie, id int) error
	Delete(ctx context.Context, id int64) error
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
	SAVE_MOVIE = "INSERT INTO movies (title, rating, awards, length, genre_id) VALUES (?,?,?,?,?);"

	GET_ALL_MOVIES = "SELECT m.id ,m.title, m.rating, m.awards, m.length, m.genre_id FROM movies m;"

	GET_ALL_MOVIES_BY_GENRE = "SELECT m.id ,m.title, m.rating, m.awards, m.length, m.genre_id, g.name FROM movies m INNER JOIN genres g ON m.genre_id = g.id WHERE g.id=?;"

	GET_MOVIE = "SELECT id, title, rating, awards, length, genre_id FROM movies WHERE id=?;"

	UPDATE_MOVIE = "UPDATE movies SET title=?, rating=?, awards=?, length=?, genre_id=? WHERE id=?;"

	DELETE_MOVIE = "DELETE FROM movies WHERE id=?;"

	EXIST_MOVIE = "SELECT m.id FROM movies m WHERE m.id=?"

	GET_MOVIE_TIMEOUT = "SELECT SLEEP(20) FROM DUAL WHERE id=?;"
)

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(EXIST_MOVIE, id)
	err := rows.Scan(&id)
	return err == nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Movie, error) {
	var movies []domain.Movie
	rows, err := r.db.Query(GET_ALL_MOVIES)
	if err != nil {
		return []domain.Movie{}, err
	}

	for rows.Next() {
		var movie domain.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id)
		if err != nil {
			return []domain.Movie{}, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *repository) GetAllMoviesByGenre(ctx context.Context, genreID int) ([]domain.Movie, error) {
	rows, err := r.db.Query(GET_ALL_MOVIES_BY_GENRE, genreID)
	if err != nil {
		return []domain.Movie{}, err
	}

	var movies []domain.Movie
	for rows.Next() {
		var movie domain.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id)
		if err != nil {
			return []domain.Movie{}, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *repository) GetMovieByID(ctx context.Context, id int) (domain.Movie, error) {
	row := r.db.QueryRow(GET_MOVIE, id)
	var movie domain.Movie
	if err := row.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id); err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

func (r *repository) GetMovieWithContext(ctx context.Context, id int) (domain.Movie, error) {
	//row := r.db.QueryRowContext(ctx, GET_MOVIE, id) //funcion OK
	row := r.db.QueryRowContext(ctx, GET_MOVIE_TIMEOUT, id)

	var movie domain.Movie
	if err := row.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id); err != nil {
		return domain.Movie{}, err
	}

	return movie, nil
}

func (r *repository) Save(ctx context.Context, m domain.Movie) (int64, error) {
	stm, err := r.db.Prepare(SAVE_MOVIE) //preparamos la consulta
	if err != nil {
		return 0, err
	}

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id)
	if err != nil {
		return 0, err
	}

	//obtenemos el ultimo id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Update(ctx context.Context, m domain.Movie, id int) error {
	stm, err := r.db.Prepare(UPDATE_MOVIE)
	if err != nil {
		return err
	}
	defer stm.Close() //cerramos para no perder memoria

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows")
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	stm, err := r.db.Prepare(DELETE_MOVIE)
	if err != nil {
		return err
	}
	defer stm.Close()

	result, err := stm.Exec(id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("error: no affected rows")
	}

	return nil
}
