package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhinavgupta21/go-ci-cd-project/models" // Adjust import path if needed
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupRouter initializes a Gin router and a real in-memory SQLite DB with migrations
func setupRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	gin.SetMode(gin.TestMode)

	// Initialize in-memory SQLite DB for tests
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err, "failed to open in-memory sqlite DB")

	// Migrate your model(s)
	err = db.AutoMigrate(&models.Book{})
	require.NoError(t, err, "failed to migrate Book model")

	router := gin.Default()

	RegisterRoutes(&Config{
		Router: router,
		DB:     db,
	})

	return router, db
}

func TestRoutesAreRegistered(t *testing.T) {
	router, db := setupRouter(t)

	db.Create(&models.Book{
		Title:  "Test Book",
		Author: "Test Author",
	})

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
	router, _ := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), "Welcome to the Book API")
}
