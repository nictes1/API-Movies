package handler

import (
	"api-movies/cmd/server/pkg/response"
	"api-movies/internal/domain"
	mock "api-movies/pkg/test/mocks/movie"
	"api-movies/pkg/test/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// dataset
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

type responseMovies struct {
	StatusCode	string
	Message		string
	Data		[]domain.Movie
}
type responseMovie struct {
	StatusCode	string
	Message		string
	Data		domain.Movie
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

// ------------------------------------------------------------------------
// Tests Request
func TestRequest(t *testing.T) {
	// arrange
	m := &mock.MockMoviesService{}
	s := serverMovies(m)

	// params
	samples := []struct{Method, Path, Body	string}{
		{Method: http.MethodGet, Path: "/movies/a"},
		{Method: http.MethodGet, Path: "/movies/genre/a"},
		{Method: http.MethodPatch, Path: "/movies/a"},
		{Method: http.MethodDelete, Path: "/movies/a"},
	}
	for i, ts := range samples {
		t.Run(fmt.Sprintf("invalid param %d", i), func(t *testing.T) {
			req, res := utils.ClientRequest(ts.Method, ts.Path, ts.Body)
			s.ServeHTTP(res, req)

			var rr response.Response
			err := json.Unmarshal(res.Body.Bytes(), &rr)

			// assert
			assert.NoError(t, err)
			assert.Equal(t, 400, res.Code)
			assert.Equal(t, ErrParseID.Error(), rr.Message)
		})
	}

	// body
	samples = []struct{Method, Path, Body	string}{
		{Method: http.MethodPatch, Path: "/movies/1", Body: ``},
	}
	for i, ts := range samples {
		t.Run(fmt.Sprintf("invalid body %d", i), func(t *testing.T) {
			req, res := utils.ClientRequest(ts.Method, ts.Path, ts.Body)
			s.ServeHTTP(res, req)

			var rr response.Response
			err := json.Unmarshal(res.Body.Bytes(), &rr)

			// assert
			assert.NoError(t, err)
			assert.Equal(t, 400, res.Code)
			assert.Equal(t, ErrBindRequest.Error(), rr.Message)
		})
	}
}

// Test Process
// Read
func TestGetAll(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	t.Run("Ok", func(t *testing.T) {
		req, res := utils.ClientRequest(http.MethodGet, "/movies", ``)
		s.ServeHTTP(res, req)

		var rr responseMovies
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, m.DataMock, rr.Data)		
	})

	t.Run("Fail", func(t *testing.T) {
		// arrange
		m.Error = "db error"

		req, res := utils.ClientRequest(http.MethodGet, "/movies", ``)
		s.ServeHTTP(res, req)

		var rr responseMovies
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 500, res.Code)
		assert.Equal(t, m.Error, rr.Message)
	})
}
func TestGetByID(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	t.Run("Ok", func(t *testing.T) {
		req, res := utils.ClientRequest(http.MethodGet, "/movies/1", ``)
		s.ServeHTTP(res, req)

		var rr responseMovie
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, m.DataMock[0], rr.Data)		
	})

	t.Run("Fail", func(t *testing.T) {
		// arrange
		m.Error = "db error"

		req, res := utils.ClientRequest(http.MethodGet, "/movies/1", ``)
		s.ServeHTTP(res, req)

		var rr responseMovie
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 500, res.Code)
		assert.Equal(t, m.Error, rr.Message)
	})
}
func TestGetAllByGenre(t *testing.T) {
	// arrange
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	t.Run("Ok", func(t *testing.T) {
		req, res := utils.ClientRequest(http.MethodGet, "/movies/genre/2", ``)
		s.ServeHTTP(res, req)

		var rr responseMovies
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, m.DataMock[0:1], rr.Data)
	})

	t.Run("Fail", func(t *testing.T) {
		// arrange
		m.Error = "db error"

		req, res := utils.ClientRequest(http.MethodGet, "/movies/genre/2", ``)
		s.ServeHTTP(res, req)

		var rr responseMovies
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 500, res.Code)
		assert.Equal(t, m.Error, rr.Message)
	})
}

// Write
func TestCreate(t *testing.T) {
	// assert
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	t.Run("Ok", func(t *testing.T) {
		expected := &domain.Movie{ID: 4, Title: "last item", Rating: 50, Awards: 50, Length: 50}

		req, res := utils.ClientRequest(http.MethodPost, "/movies", `{"id": 4, "title": "last item", "rating": 50, "awards": 50, "length": 50}`)
		s.ServeHTTP(res, req)

		var rr responseMovie
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 201, res.Code)
		assert.Equal(t, *expected, rr.Data)
	})

	t.Run("Fail", func(t *testing.T) {
		// arrage
		m.Error = "db "

		req, res := utils.ClientRequest(http.MethodPost, "/movies", `{"id": 4, "title": "last item", "rating": 50, "awards": 50, "length": 50}`)
		s.ServeHTTP(res, req)

		var rr responseMovie
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 500, res.Code)
		assert.Equal(t, m.Error, rr.Message)
	})
}
func TestUpdate(t *testing.T) {
	// assert
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	t.Run("Ok", func(t *testing.T) {
		expected := data[2]
		expected.Rating = 99; expected.Awards = 99

		req, res := utils.ClientRequest(http.MethodPatch, "/movies/3", `{"rating": 99, "awards": 99}`)
		s.ServeHTTP(res, req)

		var rr responseMovie
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expected, rr.Data)
	})

	t.Run("Fail", func(t *testing.T) {
		// arrage
		m.Error = "db error"

		expected := data[2]
		expected.Rating = 99; expected.Awards = 99

		req, res := utils.ClientRequest(http.MethodPatch, "/movies/3", `{"rating": 99, "awards": 99}`)
		s.ServeHTTP(res, req)

		var rr responseMovie
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 500, res.Code)
		assert.Equal(t, m.Error, rr.Message)
	})
}
func TestDelete(t *testing.T) {
	// assert
	data := append([]domain.Movie{}, movie_test...)
	m := &mock.MockMoviesService{
		DataMock: data,
	}
	s := serverMovies(m)

	// act
	t.Run("Ok", func(t *testing.T) {
		req, res := utils.ClientRequest(http.MethodDelete, "/movies/1", ``)
		s.ServeHTTP(res, req)

		// assert
		assert.Equal(t, 204, res.Code)
	})

	t.Run("Fail", func(t *testing.T) {
		// arrage
		m.Error = "db error"

		req, res := utils.ClientRequest(http.MethodDelete, "/movies/1", ``)
		s.ServeHTTP(res, req)

		var rr responseMovie
		err := json.Unmarshal(res.Body.Bytes(), &rr)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 500, res.Code)
		assert.Equal(t, m.Error, rr.Message)
	})
}