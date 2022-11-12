package handler

import (
	"api-movies/cmd/server/pkg/response"
	"api-movies/internal/domain"
	"api-movies/internal/movie"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	service movie.Service
}

func NewMovie(service movie.Service) *Movie {
	return &Movie{
		service: service,
	}
}

func (m *Movie) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		movies, err := m.service.GetAll(ctx)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, err)
			return
		}

		// process
		// ...

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to get movies", movies)
	}
}

func (m *Movie) GetGetAllMoviesByGenre() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		genre_id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, err)
			return
		}

		// process
		movies, err := m.service.GetAllMoviesByGenre(ctx, genre_id)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, err)
			return
		}

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to get movies by genre", movies)
	}
}

func (m *Movie) GetMovieByID() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, err)
			return
		}

		// process
		movie, err := m.service.GetMovieByID(ctx, id)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, err)
			return
		}

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to get movie by id", movie)
	}
}

func (m *Movie) GetMovieWithContext() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, err)
			return
		}

		// process
		movie, err := m.service.GetMovieWithContext(ctx, id)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, err)
			return
		}

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to get movie by id with context", movie)
	}
}

func (m *Movie) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		var movie domain.Movie
		err := ctx.ShouldBindJSON(&movie)
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, err)
			return
		}

		// process
		movie, err = m.service.Save(ctx, movie)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, err)
			return
		}

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to create movie " + movie.Title, movie)
	}
}

func (m *Movie) Update() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, err)
			return
		}
		var movie domain.Movie
		err = ctx.ShouldBindJSON(&movie)
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, err)
			return
		}

		// process
		movie, err = m.service.Update(ctx, movie, id)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, err)
			return
		}
		movie.ID = id

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to update movie", movie)
	}
}

func (m *Movie) Delete() gin.HandlerFunc {
	
	return func(ctx *gin.Context) {
		// request
		id, err := strconv.ParseInt((ctx.Param("id")), 10, 64)
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, err)
			return
		}
		
		// process
		err = m.service.Delete(ctx, id)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, err)
			return
		}

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to delete movie", nil)
	}
}
