package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/Beretta350/golang-rest-template/config"
)

// InitDatabase initializes the database connection
func InitDatabase(cfg *config.DatabaseConfig) (*sql.DB, error) {
	var sqlDB *sql.DB
	var err error

	switch cfg.Type {
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
		sqlDB, err = sql.Open("postgres", dsn)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}

	log.Printf("Establishing connection with %v database\n", cfg.Type)
	if err != nil {
		return nil, err
	}

	// Optionally, add ping and other validations
	if sqlDB != nil {
		if err = sqlDB.Ping(); err != nil {
			return nil, err
		}
	}

	return sqlDB, nil
}
