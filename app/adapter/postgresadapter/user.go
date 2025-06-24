package postgresadapter

import (
	"time"

	"github.com/Acova/movie-collection/app/domain"
	"gorm.io/gorm"
)

type PostgresUser struct {
	gorm.Model
	ID           uint
	Email        string
	Name         string
	Password     string
	DisabledDate time.Time
}

func (PostgresUser) TableName() string {
	return "user"
}

func (u *PostgresUser) ToDomain() domain.User {
	return domain.User{
		Email:        u.Email,
		Name:         u.Name,
		Password:     u.Password,
		RegisterDate: u.CreatedAt,
		DisableDate:  u.DisabledDate,
	}
}

type PostgresUserRepository struct {
	postgres *PostgresDBConnection
}

func NewPostgresUserRepository(postgres *PostgresDBConnection) *PostgresUserRepository {
	return &PostgresUserRepository{
		postgres: postgres,
	}
}

func (repository *PostgresUserRepository) ListUsers() []domain.User {
	return make([]domain.User, 0)
}

func (repository *PostgresUserRepository) CreateUser(user domain.User) {
	postgresUser := PostgresUser{
		Email:        user.Email,
		Name:         user.Name,
		Password:     user.Password,
		DisabledDate: user.DisableDate,
	}

	result := repository.postgres.DB.Create(&postgresUser)
	if result.Error != nil {
		panic("failed to create user: " + result.Error.Error())
	}
}
