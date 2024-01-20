package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ciscapello/api-service/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{
		pool: pool,
	}
}

func (r *UsersRepo) CheckUserIsExists(email string) bool {
	sqlStatement := `
	SELECT id FROM users WHERE email=$1 LIMIT 1
	`
	var id int
	err := r.pool.QueryRow(context.Background(), sqlStatement, email).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)
	return id != 0
}

func (r *UsersRepo) CreateUser(email string, password string) (int, string) {
	sqlStatement := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	RETURNING id`

	var id int
	err := r.pool.QueryRow(context.Background(), sqlStatement, email, password).Scan(&id)
	fmt.Println(err)
	if err != nil {
		fmt.Printf("error %v", err)
		return 0, "cannot create user in db"
	}
	return id, ""
}

func (r *UsersRepo) CreateSession(userId int, token string, expired_at time.Time) int {
	sqlStatement := `
	INSERT INTO sessions (user_id, refresh_token, expired_at)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	var id int
	err := r.pool.QueryRow(context.TODO(), sqlStatement, userId, token, expired_at).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
	return id
}

func (r *UsersRepo) GetByEmail(email string) (model.User, error) {
	sqlStatement := `
	SELECT id, email, password FROM users
	WHERE email=$1
	LIMIT 1
	`
	var (
		Id       int
		Email    string
		Password string
	)
	err := r.pool.QueryRow(context.Background(), sqlStatement, email).Scan(&Id, &Email, &Password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("user", Id, Email)
	user := model.User{Id: Id, Email: Email, Password: Password}
	return user, nil
}
