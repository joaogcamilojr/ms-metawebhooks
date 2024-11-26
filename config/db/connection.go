package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect () (*sql.DB, error) {
  host := os.Getenv("POSTGRES_HOST")
  port := os.Getenv("POSTGRES_PORT")
  user := os.Getenv("POSTGRES_USER")
  password := os.Getenv("POSTGRES_PASSWORD")
  database := os.Getenv("POSTGRES_DATABASE")

  if host == "" || port == "" || user == "" || password == "" || database == "" {
    log.Fatal("Check you're credentials on environment file")
  }

  conns := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable", host, port, user, password, database)

  db, err := sql.Open("postgres", conns)

  if err != nil {
    panic(err)
  }

  err = db.Ping()

  return db, err
}
