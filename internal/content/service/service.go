package service

import (
	"github.com/something-to-start-with/api-server-go/internal/content"
	"github.com/something-to-start-with/api-server-go/internal/models"
)

var _ content.Service = (*Service)(nil)

type Service struct {
	r content.Repository
}

func New(r content.Repository) *Service {
	return &Service{r: r}
}

func (s *Service) GetAll() ([]*models.Content, error) {
	return s.r.SelectAll()
}

func (s *Service) Create(c *models.Content) (*models.Content, error) {
	return s.r.Save(c)
}

func (s *Service) Update(id int, c *models.Content) (*models.Content, error) {
	err := s.r.Update(id, c)
	if err != nil {
		return nil, err
	}
	return s.r.SelectByID(id)
}

func (s *Service) GetByID(id int) (*models.Content, error) {
	return s.r.SelectByID(id)
}

func (s *Service) DeleteByID(id int) error {
	return s.r.DeleteByID(id)
}
