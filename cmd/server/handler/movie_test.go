package handler

import (
	"api-movies/internal/domain"
	mock "api-movies/pkg/test/mocks/movie"
	"api-movies/pkg/test/utils"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var movie_test = []domain.Movie{
	{
		ID:           1,
		// Created_at:   time.Now(),
		// Updated_at:   time.Now(),
		Title:        "El encargado",
		Rating:       7,
		Awards:       3,
		Release_date: time.Layout,
		Length:       180,
		Genre_id:     2,
	},
	{
		ID:           2,
		// Created_at:   time.Now(),
		// Updated_at:   time.Now(),
		Title:        "Las tortugas Ninjas",
		Rating:       9,
		Awards:       22,
		Release_date: time.Layout,
		Length:       156,
		Genre_id:     3,
	},
	{
		ID:           3,
		// Created_at:   time.Now(),
		// Updated_at:   time.Now(),
		Title:        "Rocky 2",
		Rating:       7,
		Awards:       9,
		Release_date: time.Layout,
		Length:       190,
		Genre_id:     5,
	},
}

// utils
func serverMovies(m *mock.MockMoviesService) *gin.Engine {
	// controller
	h := NewMovie(m)
	
	// server
	router := gin.Default()
	router.GET("/movies", h.GetAll())
	router.GET("/movies/genre/:id", h.GetAllMoviesByGenre())
	router.GET("/movies/:id", h.GetMovieByID())
	router.POST("/movies", h.Create())
	router.DELETE("/movies/:id", h.Delete())
	router.PATCH("/movies/:id", h.Update())

	return router
}

// Read
func TestGetAllOk(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodGet, "/movies", ``)
	s.ServeHTTP(res, req)
	
	rr := []domain.Movie{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, m.DataMock, rr)
}
func TestGetAllFail(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
		Error: "db error",
	}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodGet, "/movies", ``)
	s.ServeHTTP(res, req)
	
	rr := struct{Error string `json:"error"`}{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 500, res.Code)
	assert.Equal(t, m.Error, rr.Error)	
}

func TestGetByIDOk(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodGet, "/movies/1", ``)
	s.ServeHTTP(res, req)
	
	rr := domain.Movie{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, m.DataMock[0], rr)	
}
func TestGetByIDFail(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
		Error: "db error",
	}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodGet, "/movies/1", ``)
	s.ServeHTTP(res, req)
	
	rr := struct{Error string `json:"error"`}{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 500, res.Code)
	assert.Equal(t, m.Error, rr.Error)
}
func TestGetByIDFailController(t *testing.T) {
	// arrange
	m := &mock.MockMoviesService{}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodGet, "/movies/a", ``)
	s.ServeHTTP(res, req)
	
	rr := struct{Error string `json:"error"`}{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 400, res.Code)
	assert.NotEmpty(t, rr.Error)
}

func TestGetByGenreOk(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodGet, "/movies/genre/2", ``)
	s.ServeHTTP(res, req)
	
	rr := []domain.Movie{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, m.DataMock[0:1], rr)	
}
func TestGetByGenreFail(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
		Error: "db error",
	}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodGet, "/movies/genre/2", ``)
	s.ServeHTTP(res, req)
	
	rr := struct{Error string `json:"error"`}{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 500, res.Code)
	assert.Equal(t, m.Error, rr.Error)
}
func TestGetByGenreFailController(t *testing.T) {
	// arrange
	m := &mock.MockMoviesService{}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodGet, "/movies/genre/a", ``)
	s.ServeHTTP(res, req)
	
	rr := struct{Error string `json:"error"`}{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 400, res.Code)
	assert.NotEmpty(t, rr.Error)
}


// Write
func TestUpdateOk(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	update := data[2]
	update.Rating = 99; update.Awards = 99

	// act
	req, res := utils.ClientRequest(http.MethodPatch, "/movies/3", `{"rating": 99, "awards": 99}`)
	s.ServeHTTP(res, req)
	
	rr := struct{Movie domain.Movie `json:"movie"`}{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, update, rr.Movie)
}
func TestUpdateFail(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
		Error: "db error",
	}
	s := serverMovies(m)

	// update := data[2]
	// update.Rating = 99; update.Awards = 99

	// act
	req, res := utils.ClientRequest(http.MethodPatch, "/movies/3", `{"rating": 99, "awards": 99}`)
	s.ServeHTTP(res, req)
	
	rr := struct{Error string `json:"error"`}{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, m.Error, rr.Error)
}
func TestUpdateFailController(t *testing.T) {
	// arrange
	m := &mock.MockMoviesService{}
	s := serverMovies(m)

	samples := []struct{
		Name, Path, Body string
		Error 			 string
		StatusCode 		 int
	}{
		{Name: "fail_invalidID", Path: "/movies/a", Body: "{}", Error: "invalid ID", StatusCode: 404},
		{Name: "fail_binding", Path: "/movies/3", Body: "", Error: "EOF", StatusCode: 400},
	}

	// act
	for _, ts := range samples {
		t.Run(ts.Name, func(t *testing.T) {
			req, res := utils.ClientRequest(http.MethodPatch, ts.Path, ts.Body)
			s.ServeHTTP(res, req)

			rr := struct{Error string `json:"error"`}{}
			err:= json.Unmarshal(res.Body.Bytes(), &rr)

			// assert
			assert.NoError(t, err)
			assert.Equal(t, ts.StatusCode, res.Code)
			assert.Equal(t, ts.Error, rr.Error)
		})
	} 
}

func TestDeleteOk(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodDelete, "/movies/3", ``)
	s.ServeHTTP(res, req)

	// assert
	assert.Equal(t, 204, res.Code)
}
func TestDeleteFail(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
		Error: "db error",
	}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodDelete, "/movies/3", ``)
	s.ServeHTTP(res, req)

	rr := struct{Error string `json:"error"`}{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 500, res.Code)
	assert.Equal(t, m.Error, rr.Error)
}
func TestDeleteFailController(t *testing.T) {
	// arrange
	m := &mock.MockMoviesService{}
	s := serverMovies(m)

	// act
	req, res := utils.ClientRequest(http.MethodDelete, "/movies/a", ``)
	s.ServeHTTP(res, req)

	rr := struct{Error string `json:"error"`}{}
	err:= json.Unmarshal(res.Body.Bytes(), &rr)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "invalid ID", rr.Error)
}