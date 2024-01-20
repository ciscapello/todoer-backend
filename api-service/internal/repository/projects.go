package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	errCreate = errors.New("cannot create new project")
)

type ProjectsRepo struct {
	pool *pgxpool.Pool
}

func NewProjectRepo(pool *pgxpool.Pool) *ProjectsRepo {
	return &ProjectsRepo{
		pool: pool,
	}
}

func (r *UsersRepo) CreateProject(email string, password string, userId int) (int, error) {
	sqlStatement := `
	INSERT INTO projects (title, description, user_id)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	var id int
	err := r.pool.QueryRow(context.Background(), sqlStatement, email, password).Scan(&id)
	if err != nil {
		fmt.Printf("error %v", err)
		return 0, errCreate
	}
	return id, nil
}
