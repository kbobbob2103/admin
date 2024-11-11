package memory

import (
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
		return errors.New("สร้างข้อมูลไม่สำเร็จ")
	}
	return nil
}
