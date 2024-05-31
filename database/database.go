package database

import (
	"trainee3/database/mongo"
	"trainee3/database/mysql"

	"github.com/redis/go-redis/v9"
)

type Database struct {
	Mysql mysql.IMysqlDB
	Mongo mongo.IMongoDB
	Redis *redis.Client
}

func New(mysql mysql.IMysqlDB, mongo mongo.IMongoDB, redis *redis.Client) *Database {
	return &Database{
		Mysql: mysql,
		Mongo: mongo,
		Redis: redis,
	}
}
