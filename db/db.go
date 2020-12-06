package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/subosito/gotenv"
	"io/ioutil"
	"os"
)

var Conn *pgx.Conn
func InitConnection() {
	var err error
	_ = gotenv.Load()
	dbUrl := os.Getenv("DATABASE_URL")
	Conn, err = pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	file, err := ioutil.ReadFile("./db/db.sql")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to read db.sql: %v\n", err)
		os.Exit(1)
	}
	_, _ = Conn.Exec(context.Background(), string(file))
}
