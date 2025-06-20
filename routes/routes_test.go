package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// setupRouter creates a Gin router with nil mock DB
func setupRouter(t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	RegisterRoutes(&Config{
		Router: router,
		DB:     &gorm.DB{},
	})

	return router
}

func TestRoutesAreRegistered(t *testing.T) {
	router := setupRouter(t)

	tests := []struct {
		method string
		path   string
	}{
		{"GET", "/"},
		{"GET", "/books"},
		{"POST", "/books"},
		{"GET", "/books/1"},
		{"PUT", "/books/1"},
		{"DELETE", "/books/1"},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.method, tt.path, nil)
		router.ServeHTTP(w, req)

		require.NotEqual(t, http.StatusNotFound, w.Code, tt.method+" "+tt.path+" should be registered")
	}
}

func TestRootRouteReturnsWelcomeMessage(t *testing.T) {
	router := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), "Welcome to the Book API")
}
