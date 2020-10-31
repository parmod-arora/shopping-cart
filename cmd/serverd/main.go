package main

import (
	"cinemo.com/shoping-cart/cmd/serverd/appenv"
	"cinemo.com/shoping-cart/internal/db"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Inside main file")
	appenv.Validate()
	dbConnPool := db.InitDatabase(appenv.GetWithDefault("DATABASE_URL", "postgres://shopingcart:@localhost:5433/shopingcart?sslmode=disable"))

	log.Info(dbConnPool)
}
