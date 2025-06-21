package movies

import "github.com/Acova/movie-collection/app/model/users"

type movie struct {
	Name    string
	Creator users.User
}
