package postgresadapter

import (
	"time"

	"github.com/Acova/movie-collection/app/domain"
	"gorm.io/gorm"
)

type PostgresAdapterUser struct {
	gorm.Model
	ID           uint
	Email        string
	Name         string
	Password     string
	DisabledDate time.Time
}

func (PostgresAdapterUser) TableName() string {
	return "user"
}

func (u *PostgresAdapterUser) ToDomain() domain.User {
	return domain.User{
		Email:        u.Email,
		Name:         u.Name,
		Password:     u.Password,
		RegisterDate: u.CreatedAt,
		DisableDate:  u.DisabledDate,
	}
}

type PostgresAdapterUserRepository struct {
	postgres *PostgresDBConnection
}

func NewPostgresAdapterUserRepository(postgres *PostgresDBConnection) *PostgresAdapterUserRepository {
	return &PostgresAdapterUserRepository{
		postgres: postgres,
	}
}

func (repository *PostgresAdapterUserRepository) ListUsers() []domain.User {
	return make([]domain.User, 0)
}

func (repository *PostgresAdapterUserRepository) CreateUser(user domain.User) {
	postgresUser := PostgresAdapterUser{
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
