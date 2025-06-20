package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abhinavgupta21/go-ci-cd-project/config"
	"github.com/abhinavgupta21/go-ci-cd-project/db"
	"github.com/abhinavgupta21/go-ci-cd-project/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Using system environment variables.")
	}

	// Load all env variables into the project
	config.Initialize()

	gormDB, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Database initialization error: %v", err)
	}

	if err := db.UpdateTables(gormDB); err != nil {
		log.Fatalf("Database migration error: %v", err)
	}

	router := gin.Default()
	routes.RegisterRoutes(&routes.Config{
		Router: router,
		DB:     gormDB,
	})

	srv := startServer(router)

	gracefulShutdown(srv, gormDB)
}

func startServer(router *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":" + config.PORT,
		Handler: router,
	}

	log.Printf("Server started on: http://localhost:%s", config.PORT)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen failed: %v\n", err)
		}
	}()

	return srv
}

func gracefulShutdown(srv *http.Server, gormDB *gorm.DB) {
	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown signal received")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server stopped successfully")

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("failed to get underlying SQL DB: %v", err)
	}

	// Ensure DB is closed
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing DB: %v", err)
	}
	log.Println("Database connection closed")
}
