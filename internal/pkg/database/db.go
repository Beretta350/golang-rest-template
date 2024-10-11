package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Beretta350/golang-rest-template/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitDatabase initializes the database connection
func InitDatabase(cfg *config.DatabaseConfig) (*mongo.Database, error) {
	var mongoClient *mongo.Client
	var err error

	switch cfg.Type {
	case "mongodb":
		uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
		mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}

	log.Printf("Establishing connection with %v database\n", cfg.Type)
	if err != nil {
		return nil, err
	}

	if mongoClient != nil {

		if err = mongoClient.Ping(context.Background(), nil); err != nil {
			return nil, err
		}
	}

	return mongoClient.Database(cfg.Name), nil
}
