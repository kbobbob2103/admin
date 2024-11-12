package memory

import (
	"admin/microservice/exception"
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type permissionMongoRepository struct {
	collection *mongo.Collection
}

func NewPermissionMongoRepository(
	database *mongo.Database,
) repositoty.IPermissionRepo {
	c := &permissionMongoRepository{
		collection: database.Collection("permission"),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.collection.Indexes().
		CreateOne(ctx, mongo.IndexModel{
			Keys: bson.D{{"permission_id", 1}},
		})
	if err != nil {
		log.Println(err.Error())
	}
	return c
}
func (p permissionMongoRepository) CreatePermission(data dto.Permission) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := p.collection.InsertOne(ctx, data)
	if err != nil {
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return nil
}
func (p permissionMongoRepository) FindOnePermissionByID(id string) (dto.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{
		"permission_id": id,
	}

	var result dto.Permission
	err := p.collection.FindOne(ctx, filters).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return dto.Permission{}, exception.NewAppError(
			exception.ErrCodeNotFound,
			"ไม่พบ: "+id,
		)
	} else if err != nil {
		return dto.Permission{}, exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return result, nil
}
func (p permissionMongoRepository) UpdatePermission(data dto.Permission) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.D{
		{"permission_id", data.PermissionID}}
	update := bson.D{{"$set", data}}

	_, err := p.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return nil
}
