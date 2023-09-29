package content

import "github.com/something-to-start-with/api-server-go/internal/models"

type Service interface {
	GetAll() ([]*models.Content, error)
	Create(*models.Content) (*models.Content, error)
	Update(int, *models.Content) (*models.Content, error)
	GetByID(int) (*models.Content, error)
	DeleteByID(id int) error
}
