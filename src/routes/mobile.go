package routes

import (
	"anti-jomblo-go/src/app/mobile"

	"github.com/gin-gonic/gin"

	"anti-jomblo-go/library/data"

	"github.com/jmoiron/sqlx"
)

func RegisterMobileRoutes(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine) {
	v1 := router.Group("/mobile/v1")
	{
		mobile.RegisterRoutes(db, dataManager, router, v1)
	}
}
