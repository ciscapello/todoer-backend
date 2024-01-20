package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionsRepo struct {
	pool *pgxpool.Pool
}

func NewSessionsRepo(pool *pgxpool.Pool) *SessionsRepo {
	return &SessionsRepo{
		pool: pool,
	}
}

func (r *SessionsRepo) CreateSession(userId string, refreshToken string, expired_at time.Time) {
	sqlStatement := `
	INSERT INTO sessions VALUES
	($1, $2, $3);
	`
	err := r.pool.QueryRow(context.Background(), sqlStatement, userId, refreshToken, expired_at)
	fmt.Println(err)
}
