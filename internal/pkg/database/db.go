package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Beretta350/golang-rest-template/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitDatabase initializes the database connection
func InitDatabase(cfg *config.DatabaseConfig) (*sql.DB, *mongo.Database, error) {
	var sqlDB *sql.DB
	var mongoClient *mongo.Client
	var err error

	switch cfg.Type {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		sqlDB, err = sql.Open("mysql", dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
		sqlDB, err = sql.Open("postgres", dsn)
	case "mongodb":
		uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
		mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	default:
		return nil, nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}

	log.Printf("Establishing connection with %v database\n", cfg.Type)
	if err != nil {
		return nil, nil, err
	}

	// Optionally, add ping and other validations
	if sqlDB != nil {
		if err = sqlDB.Ping(); err != nil {
			return nil, nil, err
		}
	}
	if mongoClient != nil {

		if err = mongoClient.Ping(context.Background(), nil); err != nil {
			return nil, nil, err
		}
	}

	return sqlDB, mongoClient.Database(cfg.Name), nil
}
