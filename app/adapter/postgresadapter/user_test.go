package postgresadapter

import (
	"testing"
	"time"
)

func TestPostgresUserReturnsDoaminUser(t *testing.T) {
	now := time.Now()
	postgresUser := PostgresUser{
		Email:        "test@test.es",
		Name:         "test",
		Password:     "test",
		DisabledDate: now,
	}

	domainUser := postgresUser.ToDomain()

	if domainUser.Email != "test@test.es" {
		t.Errorf("Expected user email to be %s, got %s", "test@test.es", domainUser.Email)
	}

	if domainUser.Name != "test" {
		t.Errorf("Expected user name to be %s, got %s", "test", domainUser.Name)
	}

	if domainUser.Password != "test" {
		t.Errorf("Expected user password to be %s, got %s", "test", domainUser.Password)
	}

	if domainUser.DisableDate != now {
		t.Errorf("Expected user disable date to be %s, got %s", now, domainUser.DisableDate)
	}
}
