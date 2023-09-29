package postgres

import "github.com/something-to-start-with/api-server-go/internal/models"

type contentEntity struct {
	ID      int64  `db:"id"`
	Content string `db:"data"`
}

func (c contentEntity) toModel() *models.Content {
	return &models.Content{ID: c.ID, Body: c.Content}
}

func toModels(entities []contentEntity) []*models.Content {
	contentModels := make([]*models.Content, len(entities))
	for i, entity := range entities {
		contentModels[i] = entity.toModel()
	}
	return contentModels
}
