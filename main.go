package main

import (
	"fmt"
	"log"
	"os"

	"anti-jomblo-go/configs"
	"anti-jomblo-go/databases"
	"anti-jomblo-go/library/data"
	"anti-jomblo-go/src/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Main function for start entry golang
func main() {
	gin.SetMode(gin.TestMode)

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

	fmt.Println("Server Running...")
	routes.RegisterRoutes(db, config, dataManager)
}
