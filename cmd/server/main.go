package main

import (
	"api-movies/cmd/server/routes"
	"api-movies/pkg/db"
)

func main() {
	engine, db := db.ConnectDatabase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	engine.Run(":8080")
}
