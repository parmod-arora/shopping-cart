package main

import (
	"fmt"

	"cinemo.com/shoping-cart/application"
	"cinemo.com/shoping-cart/cmd/serverd/router"
	"cinemo.com/shoping-cart/framework/appenv"
	"cinemo.com/shoping-cart/framework/db"
	"cinemo.com/shoping-cart/framework/web/server"
	"cinemo.com/shoping-cart/internal/users"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Infof("Inside main")
	appenv.Validate()
	dbConnPool := db.InitDatabase(appenv.GetWithDefault("DATABASE_URL", "postgres://shopingcart:@localhost:5433/shopingcart?sslmode=disable"))

	app := application.App{
		UserService: users.NewUserService(dbConnPool),
	}

	s := server.New(fmt.Sprintf(":%s", appenv.GetWithDefault("PORT", "3000")), router.Handler(&app))
	s.Start()
}
