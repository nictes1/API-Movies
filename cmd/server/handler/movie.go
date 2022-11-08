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
		movies, err := m.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, movies)
	}
}

func (m *Movie) Get() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.Atoi((ctx.Param("id")))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		movie, err := m.service.Get(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, movie)
	}
}

func (m *Movie) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var movie domain.Movie
		err := ctx.ShouldBindJSON(&movie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		movie, err = m.service.Save(ctx, movie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"movie": movie.Title + " added"})
	}
}

func (m *Movie) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		movie, err = m.service.Update(ctx, movie, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		movie.ID = id
		ctx.JSON(http.StatusOK, gin.H{"movie": movie})
	}
}

func (m *Movie) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt((ctx.Param("id")), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = m.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{"delete": id})
	}
}
