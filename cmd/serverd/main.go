package main

import (
	"fmt"

	"cinemo.com/shoping-cart/application"
	"cinemo.com/shoping-cart/cmd/serverd/router"
	"cinemo.com/shoping-cart/framework/appenv"
	"cinemo.com/shoping-cart/framework/db"
	"cinemo.com/shoping-cart/framework/web/server"
	"cinemo.com/shoping-cart/internal/users"
	"cinemo.com/shoping-cart/pkg/projectpath"
	"cinemo.com/shoping-cart/pkg/yaml"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Infof("Inside main")
	// ideally this will be injected by some screct service i.e kubernates-secrect or vault
	// load jwt cert config
	vars, err := yaml.FetchEnvVarsFromYaml(projectpath.Root + "jwt-cert.yml")
	if err != nil {
		logrus.Fatalf("Error while loading jwt cert")
	}
	yaml.SetEnvVars(vars)

	// validate requied env needed for application
	appenv.Validate()

	// create database connection
	dbConnPool := db.InitDatabase(appenv.GetWithDefault("DATABASE_URL", "postgres://shopingcart:@localhost:5433/shopingcart?sslmode=disable"))

	// prepare application dependencies
	app := application.App{
		UserService: users.NewUserService(dbConnPool),
	}

	// start http server
	s := server.New(fmt.Sprintf(":%s", appenv.GetWithDefault("PORT", "3000")), router.Handler(&app))
	s.Start()
}
