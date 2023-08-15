package dbrepo

import "github.com/mjaliz/gotracktime/internal/models"

func (p *postgresDBRepo) InsertProject(pr models.ProjectInput) (models.Project, error) {
	projectDB := models.Project{
		Title: pr.Title,
	}
	if err := p.DB.Create(&projectDB).Error; err != nil {
		return models.Project{}, err
	}
	return projectDB, nil
}
