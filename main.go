package main

import (
	"flag"
	"log"
	"os"
	"time"

	"anti-jomblo-go/configs"
	"anti-jomblo-go/databases"
	"anti-jomblo-go/library/data"
	"anti-jomblo-go/src/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"
)

var loc *time.Location

type stackTracer interface {
	StackTrace() errors.StackTrace
}

var addr = flag.String("addr", ":8080", "http service address")

// Init function for initialize config
func init() {

}

// Main function for start entry golang
func main() {
	os.Setenv("TZ", "Asia/Jakarta")

	config, err := configs.GetConfiguration()
	if err != nil {
		log.Fatalln("failed to get configuration: ", err)
	}

	configs.AppConfig = config

	db, err := sqlx.Open("mysql", config.DBConnectionString)
	if err != nil {
		log.Fatalln("failed to open database x: ", err)
	}
	defer db.Close()

	dataManager := data.NewManager(
		db,
	)

	databases.MigrateUp()

	if config.ActiveWorker == 1 {
		// worker here
	}

	routes.RegisterRoutes(db, config, dataManager)
}
