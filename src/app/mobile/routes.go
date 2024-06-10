package mobile

import (
	http_user "anti-jomblo-go/src/app/mobile/user"

	"anti-jomblo-go/library/data"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	userandler http_user.UserHandler
)

func RegisterRoutes(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine, v *gin.RouterGroup) {
	v1 := v.Group("")
	{
		userandler.RegisterAPI(db, dataManager, router, v1)
	}
}
