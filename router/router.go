package router

import (
	"trainee3/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/tom-riddle-sr/logger/middleware/logger"
)

func Set(handlers *handlers.Handlers) {
	// aa := repository.Repo{
	// 	Mysql: repository.NewMysql(),
	// }

	// cfg, _ := config.New()
	// mysqlDB, _ := mysql.New(cfg)
	// db := database.Database{
	// 	Mysql: mysqlDB,
	// }.Mysql.GetDB("trainee3")

	// if err := aa.Mysql.Update(db, "id = ?", 1, &model.DSActivity{
	// 	ID:   1,
	// 	Name: "test",
	// 	Open: false,
	// }); err != nil {
	// 	logrus.Fatalf("Update error: %v", err)
	// 	fmt.Print("Update error")
	// } else {
	// 	logrus.Info("Update success")
	// 	fmt.Print("Update success") // 移動這行
	// }
	app := fiber.New(fiber.Config{
		ErrorHandler: nil,
	})

	app.Use(logger.GetLogger())
	app.Post("/auth/signIn", handlers.Auth.SignIn)
	app.Get("/cache/refresh", handlers.Cache.Refresh)

	app.Listen(":3010")

}
