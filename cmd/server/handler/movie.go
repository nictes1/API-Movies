package handler

import (
	"api-movies/cmd/server/pkg/response"
	"api-movies/internal/domain"
	"api-movies/internal/movie"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrParseID 		= errors.New("failed to parse id")
	ErrBindRequest 	= errors.New("failed to bind request")
)

type Movie struct {
	service movie.Service
}

func NewMovie(service movie.Service) *Movie {
	return &Movie{
		service: service,
	}
}

// read
func (m *Movie) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// ...

		// process
		movies, err := m.service.GetAll(ctx)
		if err != nil {
			response.Err(ctx, http.StatusInternalServerError, err)
			return
		}

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to get all movies", movies)
	}
}

func (m *Movie) GetAllMoviesByGenre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		genre_id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			response.Err(ctx, http.StatusBadRequest, ErrParseID)
			return
		}

		// process
		movies, err := m.service.GetAllMoviesByGenre(ctx, genre_id)
		if err != nil {
			response.Err(ctx, http.StatusInternalServerError, err)
			return
		}

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to get all movies by genre", movies)
	}
}

func (m *Movie) GetMovieByID() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			response.Err(ctx, http.StatusBadRequest, ErrParseID)
			return
		}

		// process
		movie, err := m.service.GetMovieByID(ctx, id)
		if err != nil {
			response.Err(ctx, http.StatusInternalServerError, err)
			return
		}

		// response
		response.Ok(ctx, http.StatusOK, "Succeed to get movie by id", movie)
	}
}

/*
func (m *Movie) GetMovieWithContext() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		movie, err := m.service.GetMovieWithContext(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, movie)
	}
}
*/

// write
func (m *Movie) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		var movie domain.Movie
		err := ctx.ShouldBindJSON(&movie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		// process
		movie, err = m.service.Save(ctx, movie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// response
		ctx.JSON(http.StatusOK, gin.H{"movie": movie.Title + " added"})
	}
}

func (m *Movie) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.Err(ctx, http.StatusBadRequest, ErrParseID)
			return
		}
		var movie domain.Movie
		err = ctx.ShouldBindJSON(&movie)
		if err != nil {
			response.Err(ctx, http.StatusBadRequest, ErrBindRequest)
			return
		}

		// process
		movie, err = m.service.Update(ctx, movie, id)
		if err != nil {
			response.Err(ctx, http.StatusInternalServerError, err)
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
			response.Err(ctx, http.StatusBadRequest, ErrParseID)
			return
		}

		// process
		err = m.service.Delete(ctx, id)
		if err != nil {
			response.Err(ctx, http.StatusInternalServerError, err)
			return
		}

		// response
		response.Ok(ctx, http.StatusNoContent, "", nil)
	}
}