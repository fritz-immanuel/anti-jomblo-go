package routes

import (
	"time"

	"github.com/gin-gonic/gin"

	"anti-jomblo-go/configs"
	"anti-jomblo-go/library/data"

	"github.com/gin-contrib/cors"
	"github.com/jmoiron/sqlx"
)

// RegisterRoutes is a base function to register all routes (api and web)
func RegisterRoutes(db *sqlx.DB, config *configs.Config, dataManager *data.Manager) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Accept", "Accept-Language", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com" //change to config
		},
		MaxAge: 12 * time.Hour, //change to config
	}))

	RegisterMobileRoutes(db, dataManager, router)

	serverAddress := config.PortApps
	router.Run(serverAddress)
}
