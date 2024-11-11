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

type employeeMongoRepository struct {
	collection *mongo.Collection
}

func NewEmployeeMongoRepository(
	database *mongo.Database,
) repositoty.IEmployeeRepo {
	c := &employeeMongoRepository{
		collection: database.Collection("employee"),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.collection.Indexes().
		CreateOne(ctx, mongo.IndexModel{
			Keys: bson.D{{"employee_id", 1}},
		})
	if err != nil {
		log.Println(err.Error())
	}
	return c
}

func (a employeeMongoRepository) CreateEmployee(data dto.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := a.collection.InsertOne(ctx, data)
	if err != nil {
		return errors.New("สร้างข้อมูลไม่สำเร็จ")
	}
	return nil
}
func (a employeeMongoRepository) UpdateOne(data dto.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.D{
		{"employee_id", data.EmployeeID}}
	update := bson.D{{"$set", data}}

	_, err := a.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("อัพเดทไม่สำเร็จ")
	}
	return nil
}
func (a employeeMongoRepository) UpdateStatus(id string, status bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{
		"employee_id": id,
	}
	update := bson.D{{"$set", bson.D{
		{Key: "status", Value: status},
	}}}

	_, err := a.collection.UpdateOne(ctx, filters, update)
	if err != nil {
		return errors.New("อัพเดทสถานะไม่สำเร็จ")
	}
	return nil
}
func (a employeeMongoRepository) UpdateStatusEmployee(id string, statusEmployee dto.StatusEmployee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{
		"employee_id": id,
	}
	update := bson.D{{"$set", bson.D{
		{Key: "status_employee", Value: statusEmployee},
	}}}

	_, err := a.collection.UpdateOne(ctx, filters, update)
	if err != nil {
		return errors.New("อัพเดทสถานะไม่สำเร็จ")
	}
	return nil
}
