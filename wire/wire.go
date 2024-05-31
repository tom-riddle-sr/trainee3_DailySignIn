//go:build wireinject
// +build wireinject

package wire

import (
	"trainee3/config"
	"trainee3/database"
	"trainee3/database/mongo"
	"trainee3/database/mysql"
	"trainee3/database/redis"
	"trainee3/handlers"
	"trainee3/repository"
	"trainee3/services"

	"github.com/google/wire"
)

type application struct {
	Services *services.Services
	Handlers *handlers.Handlers
}

func NewApplication(services *services.Services, handlers *handlers.Handlers) *application {
	return &application{
		Services: services,
		Handlers: handlers,
	}
}

func InitializeApplication() (*application, error) {
	wire.Build(
		config.New,
		mysql.New,
		mongo.New,
		redis.New,
		database.New,
		repository.NewMysql,
		repository.NewMongo,
		repository.NewRedis,
		repository.NewRepo,
		services.NewActivity,
		services.NewCache,
		services.New,
		handlers.NewActivity,
		handlers.NewCache,
		handlers.New,
		NewApplication,
	)
	return &application{}, nil
}
