package memory

import (
	"admin/microservice/exception"
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
	"context"
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
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return nil
}
func (n navigationMongoRepository) FindAllNavigationBar() ([]dto.NavigationBar, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{}
	employees := make([]dto.NavigationBar, 0)
	cursor, err := n.collection.Find(ctx, filters)
	if err != nil {
		return []dto.NavigationBar{}, exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		result := dto.NavigationBar{}
		if err = cursor.Decode(&result); err != nil {
			return []dto.NavigationBar{}, err
		}
		employees = append(employees, result)
	}
	return employees, nil
}
