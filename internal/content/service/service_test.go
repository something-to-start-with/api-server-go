package service

import (
	"database/sql"
	"github.com/something-to-start-with/api-server-go/internal/content"
	"github.com/something-to-start-with/api-server-go/internal/content/repository/testifymock"
	"github.com/something-to-start-with/api-server-go/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContentGetAll(t *testing.T) {
	tests := []struct {
		name        string
		dbContent   []*models.Content
		dbError     error
		wantContent []*models.Content
		wantError   error
	}{
		{
			"get all with return data",
			[]*models.Content{{ID: 1, Body: "body"}, {ID: 2, Body: "body2"}},
			nil,
			[]*models.Content{{ID: 1, Body: "body"}, {ID: 2, Body: "body2"}},
			nil,
		},
		{
			"get all with return error",
			[]*models.Content(nil),
			sql.ErrConnDone,
			[]*models.Content(nil),
			sql.ErrConnDone,
		},
	}

	for _, tc := range tests {
		repo := new(testifymock.Repository)
		service := New(repo)

		repo.On("SelectAll").Return(tc.dbContent, tc.dbError).Once()

		got, err := service.GetAll()

		assert.Equal(t, tc.wantContent, got)
		assert.ErrorIs(t, err, tc.wantError)
		repo.AssertExpectations(t)
	}
}

func TestContentCreate(t *testing.T) {
	tests := []struct {
		name         string
		inputContent *models.Content
		dbContent    *models.Content
		dbError      error
		wantContent  *models.Content
		wantError    error
	}{
		{
			"save with return data",
			&models.Content{Body: "body"},
			&models.Content{ID: 1, Body: "body"},
			nil,
			&models.Content{ID: 1, Body: "body"},
			nil,
		},
		{
			"save with return error",
			&models.Content{Body: "body"},
			(*models.Content)(nil),
			sql.ErrConnDone,
			(*models.Content)(nil),
			sql.ErrConnDone,
		},
	}

	for _, tc := range tests {
		repo := new(testifymock.Repository)
		service := New(repo)

		repo.On("Save", tc.inputContent).Return(tc.dbContent, tc.dbError).Once()

		got, err := service.Create(tc.inputContent)

		assert.Equal(t, tc.wantContent, got)
		assert.ErrorIs(t, err, tc.wantError)
		repo.AssertExpectations(t)
	}
}

func TestContentUpdate(t *testing.T) {
	tests := []struct {
		name         string
		inputId      int
		inputContent *models.Content
		dbContent    *models.Content
		updateError  error
		selectError  error
		selectTimes  int
		wantContent  *models.Content
		wantError    error
	}{
		{
			"update with return data",
			1,
			&models.Content{Body: "body2"},
			&models.Content{ID: 1, Body: "body"},
			nil,
			nil,
			1,
			&models.Content{ID: 1, Body: "body"},
			nil,
		},
		{
			"update with return error",
			1,
			&models.Content{Body: "body2"},
			(*models.Content)(nil),
			content.ErrNotFound,
			nil,
			0,
			(*models.Content)(nil),
			content.ErrNotFound,
		},
	}

	for _, tc := range tests {
		repo := new(testifymock.Repository)
		service := New(repo)

		mockCall := repo.On("Update", tc.inputId, tc.inputContent).Return(tc.updateError).Once()
		if tc.selectTimes > 0 {
			mockCall.On("SelectByID", tc.inputId).Return(tc.dbContent, tc.selectError).Times(tc.selectTimes)
		}

		got, err := service.Update(tc.inputId, tc.inputContent)

		assert.Equal(t, tc.wantContent, got)
		assert.ErrorIs(t, err, tc.wantError)
		repo.AssertExpectations(t)
	}
}

func TestContentGetByID(t *testing.T) {
	tests := []struct {
		name        string
		inputId     int
		dbContent   *models.Content
		dbError     error
		wantContent *models.Content
		wantError   error
	}{
		{
			"get by id with return data",
			1,
			&models.Content{ID: 1, Body: "body"},
			nil,
			&models.Content{ID: 1, Body: "body"},
			nil,
		},
		{
			"get by id with return error",
			1,
			(*models.Content)(nil),
			content.ErrNotFound,
			(*models.Content)(nil),
			content.ErrNotFound,
		},
	}

	for _, tc := range tests {
		repo := new(testifymock.Repository)
		service := New(repo)

		repo.On("SelectByID", tc.inputId).Return(tc.dbContent, tc.dbError).Once()

		got, err := service.GetByID(tc.inputId)

		assert.Equal(t, tc.wantContent, got)
		assert.ErrorIs(t, err, tc.wantError)
		repo.AssertExpectations(t)
	}
}

func TestContentDeleteByID(t *testing.T) {
	tests := []struct {
		name      string
		inputId   int
		dbContent *models.Content
		dbError   error
		wantError error
	}{
		{
			"delete by id with return data",
			1,
			&models.Content{ID: 1, Body: "body"},
			nil,
			nil,
		},
		{
			"get by id with return error",
			1,
			(*models.Content)(nil),
			content.ErrNotFound,
			content.ErrNotFound,
		},
	}

	for _, tc := range tests {
		repo := new(testifymock.Repository)
		service := New(repo)

		repo.On("DeleteByID", tc.inputId).Return(tc.dbError).Once()

		err := service.DeleteByID(tc.inputId)

		assert.ErrorIs(t, err, tc.wantError)
		repo.AssertExpectations(t)
	}
}
