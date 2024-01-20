package database

import (
	"context"
	"fmt"
	"os"

	"github.com/ciscapello/api-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
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
	// defer dbpool.Close()

	// db := stdlib.OpenDBFromPool(dbpool)
	// if err := goose.Up(db, "/internal/app/database/migrations"); err != nil {
	// 	fmt.Println(err)
	// }

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}
