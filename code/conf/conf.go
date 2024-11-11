package conf

import "os"

type Configuration struct {
	HostMongo   string
	DbMongoName string
}

func (c Configuration) EnvValidate() Configuration {

	c.HostMongo = os.Getenv("HOST_MONGO")
	c.DbMongoName = os.Getenv("DB_NAME")
	return c
}

var Env = Configuration{}.EnvValidate()
