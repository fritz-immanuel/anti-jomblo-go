package mobile

import (
	http_user "anti-jomblo-go/src/app/mobile/user"
	http_userswipe "anti-jomblo-go/src/app/mobile/userswipe"

	"anti-jomblo-go/library/data"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	userHandler      http_user.UserHandler
	userswipeHandler http_userswipe.UserSwipeHandler
)

func RegisterRoutes(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine, v *gin.RouterGroup) {
	v1 := v.Group("")
	{
		userHandler.RegisterAPI(db, dataManager, router, v1)
		userswipeHandler.RegisterAPI(db, dataManager, router, v1)
	}
}
