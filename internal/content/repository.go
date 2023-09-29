package content

import "github.com/something-to-start-with/api-server-go/internal/models"

type Repository interface {
	SelectAll() ([]*models.Content, error)
	Save(*models.Content) (*models.Content, error)
	Update(int, *models.Content) error
	SelectByID(int) (*models.Content, error)
	DeleteByID(int) error
}
