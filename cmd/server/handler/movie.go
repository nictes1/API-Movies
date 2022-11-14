package handler

import (
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
		// ...
		
		// process
		movies, err := m.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// response
		ctx.JSON(http.StatusOK, movies)
	}
}

func (m *Movie) GetAllMoviesByGenre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		genre_id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// process
		movies, err := m.service.GetAllMoviesByGenre(ctx, genre_id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// response
		ctx.JSON(http.StatusOK, movies)
	}
}

func (m *Movie) GetMovieByID() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// process
		movie, err := m.service.GetMovieByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// response
		ctx.JSON(http.StatusOK, movie)
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

func (m *Movie) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		var movie domain.Movie
		err := ctx.ShouldBindJSON(&movie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// process
		movie, err = m.service.Save(ctx, movie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// response
		ctx.JSON(http.StatusCreated, movie)
	}
}

func (m *Movie) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{"error": "invalid ID"})
			return
		}
		var movie domain.Movie
		err = ctx.ShouldBindJSON(&movie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// process
		movie, err = m.service.Update(ctx, movie, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		movie.ID = id

		// response
		ctx.JSON(http.StatusOK, gin.H{"movie": movie})
	}
}

func (m *Movie) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		id, err := strconv.ParseInt((ctx.Param("id")), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		// process
		err = m.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// response
		ctx.JSON(http.StatusNoContent, gin.H{"delete": id})
	}
}

func fieldsValidator(movie domain.Movie) string {
	var response string
	if movie.Title == "" {
		response += "el campo title no puede ser nulo" + "\n"
	}

	if movie.Rating >= 0 {
		response += "el campo Rating no puede ser menor a 0 " + "\n"
	}

	if movie.Awards >= 0 {
		response += " el campo Awards no puede ser menor a 0" + "\n"
	}

	if movie.Length > 0 {
		response += " el campo Length no puede ser menor a 0" + "\n"
	}

	if movie.Genre_id >= 0 {
		response += " el campo localitie no puede ser menor a 0" + "\n"
	}

	return response

}
