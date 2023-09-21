package api

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetContents(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/contents", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("GET /v1/contents returns code %s, should 200", strconv.Itoa(w.Code))
	}
}
