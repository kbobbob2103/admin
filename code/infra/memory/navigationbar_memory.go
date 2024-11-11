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

type navigationMongoRepository struct {
	collection *mongo.Collection
}

func NewNavigationMongoRepository(
	database *mongo.Database,
) repositoty.INavigationRepo {
	c := &navigationMongoRepository{
		collection: database.Collection("navigation_ber"),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.collection.Indexes().
		CreateOne(ctx, mongo.IndexModel{
			Keys: bson.D{{"navigation_ber_id", 1}},
		})
	if err != nil {
		log.Println(err.Error())
	}
	return c
}
func (n navigationMongoRepository) CreateNavigation(data dto.NavigationBar) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := n.collection.InsertOne(ctx, data)
	if err != nil {
		return errors.New("สร้างข้อมูลไม่สำเร็จ")
	}
	return nil
}
