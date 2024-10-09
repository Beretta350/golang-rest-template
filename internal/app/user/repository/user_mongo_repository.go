package repository

import (
	"context"
	"time"

	"github.com/Beretta350/golang-rest-template/internal/app/common/constants"
	"github.com/Beretta350/golang-rest-template/internal/app/common/logging"
	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"github.com/Beretta350/golang-rest-template/internal/pkg/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userMongoRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(d *mongo.Database) UserMongoRepository {
	return &userMongoRepository{collection: d.Collection("user")}
}

func (r *userMongoRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		logging.LogError(ctx, "repository", "GetAllUsers", err)
		return nil, errs.ErrFindingUsers.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
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

func (r *userMongoRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errs.ErrUserNotFound
	} else if err != nil {
		logging.LogError(ctx, "repository", "GetUserByID", err)
		return nil, errs.ErrFindingUserByID.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}

	return &user, nil
}

func (r *userMongoRepository) CreateUser(ctx context.Context, user *model.User) error {
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		logging.LogError(ctx, "repository", "CreateUser", err)
		return errs.ErrCreatingUser.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}

	return nil
}

func (r *userMongoRepository) UpdateUser(ctx context.Context, user *model.User) error {
	user.UpdateAt = time.Now()

	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{
		"username":  user.Username,
		"password":  user.Password,
		"update_at": user.UpdateAt,
	}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logging.LogError(ctx, "repository", "UpdateUser", err)
		return errs.ErrUpdatingUser.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}

	return nil
}

func (r *userMongoRepository) DeleteUser(ctx context.Context, id string) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		logging.LogError(ctx, "repository", "DeleteUser", err)
		return errs.ErrDeletingUser.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}

	if result.DeletedCount == 0 {
		return errs.ErrDeletingUser.SetDetailFromString(constants.NoUsersToDelete)
	}

	return nil
}
