package testifymock

import (
	"github.com/something-to-start-with/api-server-go/internal/content"
	"github.com/something-to-start-with/api-server-go/internal/models"
	"github.com/stretchr/testify/mock"
)

var _ content.Service = (*Service)(nil)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll() ([]*models.Content, error) {
	args := s.Called()
	return args.Get(0).([]*models.Content), args.Error(1)
}

func (s *Service) Create(c *models.Content) (*models.Content, error) {
	args := s.Called(c)
	return args.Get(0).(*models.Content), args.Error(1)
}

func (s *Service) Update(id int, c *models.Content) (*models.Content, error) {
	args := s.Called(id, c)
	return args.Get(0).(*models.Content), args.Error(1)
}

func (s *Service) GetByID(id int) (*models.Content, error) {
	args := s.Called(id)
	return args.Get(0).(*models.Content), args.Error(1)
}

func (s *Service) DeleteByID(id int) error {
	args := s.Called(id)
	return args.Error(0)
}
