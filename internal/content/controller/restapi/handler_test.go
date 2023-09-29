package restapi

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/something-to-start-with/api-server-go/internal/content"
	"github.com/something-to-start-with/api-server-go/internal/content/service/testifymock"
	"github.com/something-to-start-with/api-server-go/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetContents(t *testing.T) {
	tests := []struct {
		name           string
		serviceContent []*models.Content
		serviceError   error
		wantCode       int
		wantBody       []*contentResponse
	}{
		{
			"get all with return data",
			[]*models.Content{{ID: 1, Body: "body"}, {ID: 2, Body: "body2"}},
			nil,
			http.StatusOK,
			[]*contentResponse{{ID: 1, Body: "body"}, {ID: 2, Body: "body2"}},
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("GetAll").Return(tc.serviceContent, tc.serviceError)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/contents", nil)
		router.ServeHTTP(w, req)

		expectedBody, err := json.Marshal(tc.wantBody)
		require.NoError(t, err)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	}
}

func TestGetContentsError(t *testing.T) {
	tests := []struct {
		name           string
		serviceContent []*models.Content
		serviceError   error
		wantCode       int
		wantBody       string
	}{
		{
			"get all with return error",
			[]*models.Content(nil),
			sql.ErrConnDone,
			http.StatusInternalServerError,
			"",
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("GetAll").Return(tc.serviceContent, tc.serviceError)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/contents", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, tc.wantBody, w.Body.String())
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name           string
		body           *contentRequest
		serviceRequest *models.Content
		serviceContent *models.Content
		serviceError   error
		wantCode       int
		wantBody       *contentResponse
	}{
		{
			"create with return data",
			&contentRequest{Body: "body1"},
			&models.Content{Body: "body1"},
			&models.Content{ID: 1, Body: "body1"},
			nil,
			http.StatusCreated,
			&contentResponse{ID: 1, Body: "body1"},
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("Create", tc.serviceRequest).Return(tc.serviceContent, tc.serviceError)

		body, err := json.Marshal(tc.body)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/v1/contents", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		expectedBody, err := json.Marshal(tc.wantBody)
		require.NoError(t, err)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	}
}

func TestCreateError(t *testing.T) {
	tests := []struct {
		name           string
		body           *contentRequest
		serviceRequest *models.Content
		serviceContent *models.Content
		serviceError   error
		wantCode       int
		wantBody       string
	}{
		{
			"create with return error",
			&contentRequest{Body: "body1"},
			&models.Content{Body: "body1"},
			(*models.Content)(nil),
			sql.ErrConnDone,
			http.StatusInternalServerError,
			"",
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("Create", tc.serviceRequest).Return(tc.serviceContent, tc.serviceError)

		body, err := json.Marshal(tc.body)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/v1/contents", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, tc.wantBody, w.Body.String())
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		body           *contentRequest
		serviceRequest *models.Content
		serviceContent *models.Content
		serviceError   error
		wantCode       int
		wantBody       *contentResponse
	}{
		{
			"create with return data",
			1,
			&contentRequest{Body: "body1"},
			&models.Content{Body: "body1"},
			&models.Content{ID: 1, Body: "body1"},
			nil,
			http.StatusOK,
			&contentResponse{ID: 1, Body: "body1"},
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("Update", tc.id, tc.serviceRequest).Return(tc.serviceContent, tc.serviceError)

		body, err := json.Marshal(tc.body)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/v1/contents/"+strconv.Itoa(1), bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		expectedBody, err := json.Marshal(tc.wantBody)
		require.NoError(t, err)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	}
}

func TestUpdate400Error(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		body     *contentRequest
		wantCode int
		wantBody string
	}{
		{
			"create with return data",
			"a",
			&contentRequest{Body: "body1"},
			http.StatusBadRequest,
			"",
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		body, err := json.Marshal(tc.body)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/v1/contents/"+tc.id, bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, tc.wantBody, w.Body.String())
	}
}

func TestUpdateError(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		body           *contentRequest
		serviceRequest *models.Content
		serviceContent *models.Content
		serviceError   error
		wantCode       int
		wantBody       string
	}{
		{
			"create with return error",
			1,
			&contentRequest{Body: "body1"},
			&models.Content{Body: "body1"},
			(*models.Content)(nil),
			content.ErrNotFound,
			http.StatusNotFound,
			"",
		},
		{
			"create with return error",
			1,
			&contentRequest{Body: "body1"},
			&models.Content{Body: "body1"},
			(*models.Content)(nil),
			sql.ErrConnDone,
			http.StatusInternalServerError,
			"",
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("Update", tc.id, tc.serviceRequest).Return(tc.serviceContent, tc.serviceError)

		body, err := json.Marshal(tc.body)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/v1/contents/"+strconv.Itoa(1), bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, tc.wantBody, w.Body.String())
	}
}

func TestGetById(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		serviceContent *models.Content
		serviceError   error
		wantCode       int
		wantBody       *contentResponse
	}{
		{
			"get by id with return data",
			1,
			&models.Content{ID: 1, Body: "body1"},
			nil,
			http.StatusOK,
			&contentResponse{ID: 1, Body: "body1"},
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("GetByID", tc.id).Return(tc.serviceContent, tc.serviceError)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/contents/"+strconv.Itoa(tc.id), nil)
		router.ServeHTTP(w, req)

		expectedBody, err := json.Marshal(tc.wantBody)
		require.NoError(t, err)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	}
}

func TestGetByIdError(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		serviceContent *models.Content
		serviceError   error
		wantCode       int
		wantBody       string
	}{
		{
			"get by id with return error",
			1,
			(*models.Content)(nil),
			content.ErrNotFound,
			http.StatusNotFound,
			"",
		},
		{
			"get by id with return error",
			1,
			(*models.Content)(nil),
			sql.ErrConnDone,
			http.StatusInternalServerError,
			"",
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("GetByID", tc.id).Return(tc.serviceContent, tc.serviceError)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/contents/"+strconv.Itoa(tc.id), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, tc.wantBody, w.Body.String())
	}
}

func TestDeleteById(t *testing.T) {
	tests := []struct {
		name         string
		id           int
		serviceError error
		wantCode     int
		wantBody     string
	}{
		{
			"delete by id with return data",
			1,
			nil,
			http.StatusOK,
			"",
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("DeleteByID", tc.id).Return(tc.serviceError)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/v1/contents/"+strconv.Itoa(tc.id), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, tc.wantBody, w.Body.String())
	}
}

func TestDeleteByIdError(t *testing.T) {
	tests := []struct {
		name         string
		id           int
		serviceError error
		wantCode     int
		wantBody     string
	}{
		{
			"delete by id with return error",
			1,
			content.ErrNotFound,
			http.StatusNotFound,
			"",
		},
		{
			"delete by id with return error",
			1,
			sql.ErrConnDone,
			http.StatusInternalServerError,
			"",
		},
	}

	for _, tc := range tests {
		router := gin.Default()

		contentService := new(testifymock.Service)
		SetupRoutes(router, contentService)

		contentService.On("DeleteByID", tc.id).Return(tc.serviceError)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/v1/contents/"+strconv.Itoa(tc.id), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.wantCode, w.Code)
		assert.Equal(t, tc.wantBody, w.Body.String())
	}
}
