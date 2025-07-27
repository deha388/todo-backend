package database

import (
	"fmt"
	"log"
	"todo-backend/internal/infrastructure/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewConnection creates a new database connection
func NewConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDatabaseDSN()
	
	// Configure GORM logger
	logLevel := logger.Error
	if cfg.Logging.Level == "info" {
		logLevel = logger.Info
	}

	// Use SQLite as default database
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite database: %w", err)
	}

	// Auto-migrate the schema - use SQLiteTodoModel
	if err := db.AutoMigrate(&SQLiteTodoModel{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Printf("✅ SQLite database connected: %s", dsn)
	return db, nil
} 