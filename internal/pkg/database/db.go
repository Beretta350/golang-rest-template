package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Beretta350/golang-rest-template/config"
)

// InitDatabase initializes the database connection
func InitDatabase(cfg *config.DatabaseConfig) (*sql.DB, error) {
	var sqlDB *sql.DB
	var err error

	switch cfg.Type {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		fmt.Println("connection:", dsn)
		sqlDB, err = sql.Open("mysql", dsn)
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
