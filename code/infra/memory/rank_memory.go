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

type rankMongoRepository struct {
	collection *mongo.Collection
}

func NewRankMongoRepository(
	database *mongo.Database,
) repositoty.IRankRepo {
	c := &rankMongoRepository{
		collection: database.Collection("rank"),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.collection.Indexes().
		CreateOne(ctx, mongo.IndexModel{
			Keys: bson.D{{"rank_id", 1}},
		})
	if err != nil {
		log.Println(err.Error())
	}
	return c
}
func (r rankMongoRepository) CreateRank(data dto.Rank) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, data)
	if err != nil {
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return nil
}
func (r rankMongoRepository) FindAllRank() ([]dto.Rank, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{}
	ranks := make([]dto.Rank, 0)
	cursor, err := r.collection.Find(ctx, filters)
	if err != nil {
		return []dto.Rank{}, exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		result := dto.Rank{}
		if err = cursor.Decode(&result); err != nil {
			return []dto.Rank{}, err
		}
		ranks = append(ranks, result)
	}
	return ranks, nil
}
