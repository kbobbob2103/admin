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

type roleMongoRepository struct {
	collection *mongo.Collection
}

func NewRoleMongoRepository(
	database *mongo.Database,
) repositoty.IRoleRepo {
	c := &roleMongoRepository{
		collection: database.Collection("role"),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.collection.Indexes().
		CreateOne(ctx, mongo.IndexModel{
			Keys: bson.D{{"role_id", 1}},
		})
	if err != nil {
		log.Println(err.Error())
	}
	return c
}

func (a roleMongoRepository) CreateRole(data dto.Role) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := a.collection.InsertOne(ctx, data)
	if err != nil {
		return errors.New("สร้างข้อมูลไม่สำเร็จ")
	}
	return nil
}
