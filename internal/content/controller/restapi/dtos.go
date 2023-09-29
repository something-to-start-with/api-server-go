package restapi

import "github.com/something-to-start-with/api-server-go/internal/models"

type contentRequest struct {
	Body string `json:"body"`
}

type contentResponse struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
}

func (cr contentRequest) toModel() *models.Content {
	return &models.Content{Body: cr.Body}
}

func newContentsResponses(models []*models.Content) []*contentResponse {
	contents := make([]*contentResponse, len(models))
	for i, model := range models {
		contents[i] = newContentResponse(model)
	}
	return contents
}

func newContentResponse(m *models.Content) *contentResponse {
	return &contentResponse{ID: m.ID, Body: m.Body}
}
