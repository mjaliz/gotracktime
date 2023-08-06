package repository

import "github.com/mjaliz/gotracktime/internal/inputs"

type DatabaseRepo interface {
	InsertUser(user inputs.User) error
}
