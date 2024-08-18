package database

import (
	"fmt"
	"log"
	// "os"
	"database/sql"

	"github.com/MukuFlash03/hackernews/internal/utils"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "mukuflash"
    password = "mukuflash"
    dbname   = "hackernews"
)

var PG_DB_URL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
const MIGRATION_PATH = "file://internal/pkg/db/migrations/postgres"

var Db *sql.DB

func InitDB() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlconn)
    utils.CheckError(err, "panic")
      
    err = db.Ping()
    utils.CheckError(err, "panic")
 
	Db = db
    fmt.Println("Connected!")
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	m, err := migrate.New(
		"file://internal/pkg/db/migrations/postgres",
		PG_DB_URL,
	)

    if err != nil {
        log.Fatalf("Error creating migration instance: %v", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Error running migrations: %v", err)
    }

    log.Println("Migrations completed successfully")
}