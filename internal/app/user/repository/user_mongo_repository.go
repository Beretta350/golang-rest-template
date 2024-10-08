package repository

import (
	"context"
	"time"

	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"github.com/Beretta350/golang-rest-template/internal/pkg/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type userMongoRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(d *mongo.Database) UserMongoRepository {
	return &userMongoRepository{collection: d.Collection("user")}
}

func (r *userMongoRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errs.ErrFindingUsers.SetDetail(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		user.Password = ""
		user.Id = ""
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userMongoRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errs.ErrUserNotFound
	} else if err != nil {
		return nil, errs.ErrFindingUserByID.SetDetail(err)
	}

	return &user, nil
}

func (r *userMongoRepository) Create(ctx context.Context, user *model.User) error {
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return errs.ErrCreatingUser.SetDetail(err)
	}

	return nil
}

func (r *userMongoRepository) Update(ctx context.Context, user *model.User) error {
	user.UpdateAt = time.Now()

	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{
		"username":  user.Username,
		"password":  user.Password,
		"update_at": user.UpdateAt,
	}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errs.ErrUpdatingUser.SetDetail(err)
	}

	return nil
}

func (r *userMongoRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errs.ErrDeletingUser.SetDetail(err)
	}

	return nil
}
