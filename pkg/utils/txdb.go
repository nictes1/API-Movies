package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DATA-DOG/go-txdb"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	txdb.Register("txdb", "mysql", fmt.Sprintf("%s:%s@/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME")))
}

// instancia txdb
func InitTxDB() *sql.DB {
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
