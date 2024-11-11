package injector

import (
	"admin/microservice/conf"
	"admin/microservice/infra/memory"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var NewAdminMongoClient *mongo.Database = func() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	host, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Env.HostMongo))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return host.Database(conf.Env.DbMongoName)
}()

var (
	PermissionRepository = memory.NewPermissionMongoRepository(NewAdminMongoClient)
)
