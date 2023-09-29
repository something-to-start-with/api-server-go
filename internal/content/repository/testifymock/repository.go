package testifymock

import (
	"github.com/something-to-start-with/api-server-go/internal/content"
	"github.com/something-to-start-with/api-server-go/internal/models"
	"github.com/stretchr/testify/mock"
)

var _ content.Repository = (*Repository)(nil)

type Repository struct {
	mock.Mock
}

func (r *Repository) SelectAll() ([]*models.Content, error) {
	args := r.Called()
	return args.Get(0).([]*models.Content), args.Error(1)
}

func (r *Repository) Save(m *models.Content) (*models.Content, error) {
	args := r.Called(m)
	return args.Get(0).(*models.Content), args.Error(1)
}

func (r *Repository) Update(id int, c *models.Content) error {
	args := r.Called(id, c)
	return args.Error(0)
}

func (r *Repository) SelectByID(id int) (*models.Content, error) {
	args := r.Called(id)
	return args.Get(0).(*models.Content), args.Error(1)
}

func (r *Repository) DeleteByID(id int) error {
	args := r.Called(id)
	return args.Error(0)
}
