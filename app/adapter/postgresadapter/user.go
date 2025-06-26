package postgresadapter

import (
	"time"

	"github.com/Acova/movie-collection/app/domain"
	"gorm.io/gorm"
)

type PostgresUser struct {
	gorm.Model
	ID           uint
	Email        string    `gorm:"not null"`
	Name         string    `gorm:"not null"`
	Password     string    `gorm:"not null"`
	DisabledDate time.Time `gorm:"default:NULL"`
}

func (PostgresUser) TableName() string {
	return "user"
}

func (u *PostgresUser) ToDomain() *domain.User {
	return &domain.User{
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

func NewPostgresUserRepository(postgres *PostgresDBConnection) (*PostgresUserRepository, error) {
	return &PostgresUserRepository{
		postgres: postgres,
	}, nil
}

func (repository *PostgresUserRepository) ListUsers() ([]*domain.User, error) {
	var postgresUsers []PostgresUser
	result := repository.postgres.DB.Find(&postgresUsers)
	if result.Error != nil {
		return nil, result.Error
	}

	users := make([]*domain.User, len(postgresUsers))
	for i, postgresUser := range postgresUsers {
		users[i] = postgresUser.ToDomain()
	}

	return users, nil
}

func (repository *PostgresUserRepository) CreateUser(user *domain.User) error {
	postgresUser := PostgresUser{
		Email:        user.Email,
		Name:         user.Name,
		Password:     user.Password,
		DisabledDate: user.DisableDate,
	}

	result := repository.postgres.DB.Create(&postgresUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *PostgresUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var postgresUser PostgresUser
	result := repository.postgres.DB.Where("email = ?", email).First(&postgresUser)
	if result.Error != nil {
		return &domain.User{}, result.Error
	}

	return postgresUser.ToDomain(), nil
}
