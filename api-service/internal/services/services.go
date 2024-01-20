package services

import (
	"github.com/ciscapello/api-service/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Services struct {
	UserService    *UserService
	ProjectService *ProjectService
}

func Init(pool *pgxpool.Pool) *Services {
	usersRepo := repository.NewUserRepo(pool)
	sessionsRepo := repository.NewSessionsRepo(pool)
	projectsRepo := repository.NewProjectRepo(pool)
	return &Services{
		UserService:    NewUserService(usersRepo, sessionsRepo),
		ProjectService: NewProjectService(projectsRepo, pool),
	}
}
