package testutil

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cinemo.com/shoping-cart/framework/db"
	"cinemo.com/shoping-cart/pkg/projectpath"
	"cinemo.com/shoping-cart/pkg/trace"

	// include migrate file driver
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"

	// import pq driver
	_ "github.com/lib/pq"
)

// PrepareDatabase helps create unique schema database
func PrepareDatabase(traceInfo trace.Info) (*sql.DB, string, error) {

	schema := "schema_" + fmt.Sprintf("%x", md5.Sum([]byte(traceInfo.FunctionName)))
	migrateDbConnPool := db.InitDatabase(os.Getenv("DATABASE_URL"))
	defer func() {
		migrateDbConnPool.Close()
	}()
	migrateDbConnPool.Exec("DROP SCHEMA IF EXISTS " + schema + " CASCADE")
	_, err := migrateDbConnPool.Exec("CREATE SCHEMA " + schema)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	dbConnPool := db.InitDatabase(os.Getenv("DATABASE_URL") + "&search_path=" + schema)
	driver, err := postgres.WithInstance(dbConnPool, &postgres.Config{})
	if err != nil {
		log.Fatalf("=====error: %s", err.Error())
		return nil, schema, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+projectpath.Root+"/data/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("=====error: %s", err.Error())
		return nil, schema, err
	}
	m.Up()
	return dbConnPool, schema, err
}

// LoadFixture will load and execute SQL queries from fixture file
func LoadFixture(dbConnPool *sql.DB, fixturePath string) error {
	if fixturePath != "" {
		input, err := ioutil.ReadFile(fixturePath)
		if err != nil {
			return err
		}
		queries := strings.Split(string(input), ";")
		for _, query := range queries {
			_, err = dbConnPool.Exec(query)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
