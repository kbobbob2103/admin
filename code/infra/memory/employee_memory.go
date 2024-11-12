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
	"regexp"
	"strings"
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
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return nil
}
func (a employeeMongoRepository) FindAll(
	employeeIds,
	roleId []string,
	search string,
	page,
	limit int,
) ([]dto.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{
		"status_employee": 1,
	}
	if len(employeeIds) != 0 {
		filters["employee_id"] = bson.M{"$in": employeeIds}
	}
	if len(roleId) != 0 {
		filters["role_id"] = bson.M{"$in": roleId}
	}
	employees := make([]dto.Employee, 0)
	cursor, err := a.collection.Find(ctx, filters)
	if err != nil {
		return []dto.Employee{}, exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	defer cursor.Close(ctx)

	search = strings.ToUpper(search)
	regex := regexp.MustCompile(search)
	for cursor.Next(ctx) {
		result := dto.Employee{}
		if err = cursor.Decode(&result); err != nil {
			return []dto.Employee{}, err
		}
		itemName := strings.ToUpper(result.EmployeeName)
		if regex.MatchString(itemName) {
			employees = append(employees, result)
		}
	}

	if page != 0 && limit != 0 {
		return PaginateEmployee(employees, page, limit)
	}
	return employees, nil
}
func (a employeeMongoRepository) FindOneByID(id string) (dto.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{
		"employee_id": id,
	}

	var result dto.Employee
	err := a.collection.FindOne(ctx, filters).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return dto.Employee{}, exception.NewAppError(
			exception.ErrCodeNotFound,
			"ไม่พบ: "+id,
		)
	} else if err != nil {
		return dto.Employee{}, exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return result, nil
}
func (a employeeMongoRepository) UpdateOne(data dto.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.D{
		{"employee_id", data.EmployeeID}}
	update := bson.D{{"$set", data}}

	_, err := a.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
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
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
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
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return nil
}
func (a employeeMongoRepository) FindCount(
	employeeIds,
	roleId []string,
) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{}
	if len(employeeIds) != 0 {
		filters["employee_id"] = employeeIds
	}
	if len(roleId) != 0 {
		filters["role_id"] = roleId
	}
	count, err := a.collection.CountDocuments(ctx, filters)
	if err != nil {
		return 0, exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}

	return count, nil
}
