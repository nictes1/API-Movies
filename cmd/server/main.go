package main

import (
	"api-movies/cmd/server/routes"
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	db, err := sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	engine := gin.Default()

	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	engine.Run(":8080")
}
