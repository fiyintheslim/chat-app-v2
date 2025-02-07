package db_conn

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Psql *pgxpool.Pool



func ConnectDB()  {
	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"))
	ctx := context.Background()

	var err error

	Psql, err = pgxpool.New(ctx, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed with error: %v", err)
	}

	fmt.Println(url)
}
