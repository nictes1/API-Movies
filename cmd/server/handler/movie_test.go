package handler

import (
	"api-movies/cmd/server/pkg/response"
	"api-movies/internal/movie"
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// utils
func ServerDB() (db *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Loading .env")
	}

	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	db, err = sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func ServerInit() *gin.Engine {
	// internal
	d := ServerDB()
	r := movie.NewRepository(d)
	s := movie.NewService(r)
	
	// server + controller
	h := NewMovie(s)
	router := gin.Default()
	
	router.GET("/movies", h.GetAll())
	router.GET("/movies/genre/:id", h.GetGetAllMoviesByGenre())
	router.GET("/movies/:id", h.GetMovieByID())
	router.POST("/movies", h.Create())
	router.DELETE("/movies/:id", h.Delete())
	router.PATCH("/movies/:id", h.Update())

	return router
}
func ClientRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	// request + headers
	var req = httptest.NewRequest(method, url,  bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	// response
	var res = httptest.NewRecorder()

	return req, res
}


// tests
func TestGetAllSimple(t *testing.T) {
	// arrange
	server := ServerInit()
	req, res := ClientRequest(http.MethodGet, "/movies", ``)
	
	// act
	server.ServeHTTP(res, req)
	var rs response.Response
	err := json.Unmarshal(res.Body.Bytes(), &rs)
	
	// arrange
	assert.NoError(t, err)
}