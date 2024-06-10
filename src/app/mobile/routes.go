package mobile

import (
	http_premiumpackage "anti-jomblo-go/src/app/mobile/premiumpackage"
	http_user "anti-jomblo-go/src/app/mobile/user"
	http_userpremium "anti-jomblo-go/src/app/mobile/userpremium"
	http_userswipe "anti-jomblo-go/src/app/mobile/userswipe"

	"anti-jomblo-go/library/data"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	premiumpackageHandler http_premiumpackage.PremiumPackageHandler
	userHandler           http_user.UserHandler
	userpremiumHandler    http_userpremium.UserPremiumHandler
	userswipeHandler      http_userswipe.UserSwipeHandler
)

func RegisterRoutes(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine, v *gin.RouterGroup) {
	v1 := v.Group("")
	{
		premiumpackageHandler.RegisterAPI(db, dataManager, router, v1)
		userHandler.RegisterAPI(db, dataManager, router, v1)
		userpremiumHandler.RegisterAPI(db, dataManager, router, v1)
		userswipeHandler.RegisterAPI(db, dataManager, router, v1)
	}
}
