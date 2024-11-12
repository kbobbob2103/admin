package memory

import (
	"admin/microservice/exception"
	"admin/microservice/helpers"
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
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return nil
}
func (a roleMongoRepository) UpdateRole(data dto.Role) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.D{
		{"role_id", data.RoleID}}
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
func (a roleMongoRepository) FindAll(
	roleId []string,
	search string,
	page,
	limit int,
) ([]dto.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{
		"is_deleted": false,
	}
	if len(roleId) != 0 {
		filters["role_id"] = roleId
	}
	employees := make([]dto.Role, 0)
	cursor, err := a.collection.Find(ctx, filters)
	if err != nil {
		return []dto.Role{}, exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	defer cursor.Close(ctx)

	search = strings.ToUpper(search)
	regex := regexp.MustCompile(search)
	for cursor.Next(ctx) {
		result := dto.Role{}
		if err = cursor.Decode(&result); err != nil {
			return []dto.Role{}, err
		}
		itemName := strings.ToUpper(result.RoleName)
		if regex.MatchString(itemName) {
			employees = append(employees, result)
		}
	}

	if page != 0 && limit != 0 {
		return PaginateRole(employees, page, limit)
	}
	return employees, nil
}
func (a roleMongoRepository) FindOneRoleByID(id string) (dto.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filters := bson.M{
		"role_id": id,
	}

	var result dto.Role
	err := a.collection.FindOne(ctx, filters).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return dto.Role{}, exception.NewAppError(
			exception.ErrCodeNotFound,
			"ไม่พบ: "+id,
		)
	} else if err != nil {
		return dto.Role{}, exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return result, nil
}
func (a roleMongoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filters := bson.D{{"role_id", id}}
	updated := bson.D{{"$set", bson.D{
		{Key: "is_deleted", Value: true},
		{Key: "time_at.updated_at", Value: helpers.GetCurrentUnix()},
		{Key: "time_at.human_readable_updated_at", Value: helpers.TimeNowStr()},
	}}}

	_, err := a.collection.UpdateOne(ctx, filters, updated)
	if err != nil {
		return exception.NewAppError(
			exception.ErrCodeDatabase,
			err.Error(),
		)
	}
	return nil
}
