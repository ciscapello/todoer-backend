package database

import (
	"context"
	"fmt"
	"os"

	"github.com/ciscapello/api-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func Connect() *pgxpool.Pool {
	conf := config.New()

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		conf.DatabaseConfig.DbUser,
		conf.DatabaseConfig.DbPassword,
		conf.DatabaseConfig.DbHost,
		conf.DatabaseConfig.DbPort,
		conf.DatabaseConfig.DbName)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	db := stdlib.OpenDBFromPool(dbpool)

	if err := goose.Up(db, "/app/internal/app/database/migrations"); err != nil {
		fmt.Println(err)
	}

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}
