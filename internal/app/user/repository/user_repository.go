package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(d *mongo.Database) UserRepository {
	return &userRepository{collection: d.Collection("user")}
}

func (r *userRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, fmt.Errorf("error fetching user by id: %v", err)
	}

	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	user.UpdateAt = time.Now()

	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{
		"username":  user.Username,
		"password":  user.Password,
		"update_at": user.UpdateAt,
	}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}
