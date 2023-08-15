package repository

import (
	"github.com/mjaliz/gotracktime/internal/models"
)

type DatabaseRepo interface {
	InsertUser(user models.SignUpInput) (models.User, error)
	FindUserByEmail(user models.SignInInput) (models.User, error)
	InsertTimeEntity(timeEntity models.TimeEntityInput) (models.TimeEntity, error)
	InsertProject(project models.ProjectInput) (models.Project, error)
}
