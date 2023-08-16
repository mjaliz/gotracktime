package dbrepo

import "github.com/mjaliz/gotracktime/internal/models"

func (p *postgresDBRepo) InsertTimeEntity(te models.TimeEntityInput) (models.TimeEntity, error) {
	timeEntityDB := models.TimeEntity{
		UserID:        te.UserID,
		StartedAt:     te.StartedAt,
		DescriptionID: te.DescriptionID,
		ProjectID:     te.ProjectID,
	}
	if err := p.DB.Create(&timeEntityDB).Error; err != nil {
		return models.TimeEntity{}, err
	}
	return timeEntityDB, nil
}

func (p *testDBRepo) InsertTimeEntity(te models.TimeEntityInput) (models.TimeEntity, error) {
	timeEntityDB := models.TimeEntity{
		UserID:        te.UserID,
		StartedAt:     te.StartedAt,
		DescriptionID: te.DescriptionID,
		ProjectID:     te.ProjectID,
	}
	return timeEntityDB, nil
}
