package userpremium

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"anti-jomblo-go/library"
	"anti-jomblo-go/middleware"
	"anti-jomblo-go/models"
	"anti-jomblo-go/src/services/premiumpackage"
	"anti-jomblo-go/src/services/userpremium"

	"github.com/gin-gonic/gin"

	"anti-jomblo-go/library/appcontext"
	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/http/response"
	"anti-jomblo-go/library/types"

	userpremiumRepository "anti-jomblo-go/src/services/userpremium/repository"
	userpremiumUsecase "anti-jomblo-go/src/services/userpremium/usecase"

	premiumpackageRepository "anti-jomblo-go/src/services/premiumpackage/repository"
	premiumpackageUsecase "anti-jomblo-go/src/services/premiumpackage/usecase"
)

var ()

type UserPremiumHandler struct {
	UserPremiumUsecase    userpremium.Usecase
	PremiumPackageUsecase premiumpackage.Usecase
	dataManager           *data.Manager
	Result                gin.H
	Status                int
}

func (h UserPremiumHandler) RegisterAPI(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine, v *gin.RouterGroup) {
	userpremiumRepo := userpremiumRepository.NewUserPremiumRepository(
		data.NewMySQLStorage(db, "user_premium", models.UserPremium{}, data.MysqlConfig{}),
	)

	uUserPremium := userpremiumUsecase.NewUserPremiumUsecase(db, &userpremiumRepo)

	premiumpackageRepo := premiumpackageRepository.NewPremiumPackageRepository(
		data.NewMySQLStorage(db, "premium_packages", models.PremiumPackage{}, data.MysqlConfig{}),
		data.NewMySQLStorage(db, "status", models.Status{}, data.MysqlConfig{}),
	)

	uPremiumPackage := premiumpackageUsecase.NewPremiumPackageUsecase(db, &premiumpackageRepo)

	base := &UserPremiumHandler{
		UserPremiumUsecase:    uUserPremium,
		PremiumPackageUsecase: uPremiumPackage,
		dataManager:           dataManager,
	}

	rs := v.Group("/user-premium")
	{
		rs.POST("", middleware.AuthMobile, base.Create)
	}
}

func (h *UserPremiumHandler) Create(c *gin.Context) {
	var obj models.UserPremium
	var data *models.UserPremium

	now := library.UTCPlus7()

	obj.UserID = *appcontext.UserID(c)
	obj.BoughtAt = now

	premiumPackageID := c.PostForm("PremiumPackageID")

	errTransaction := h.dataManager.RunInTransaction(c, func(tctx *gin.Context) *types.Error {
		premiumPackageData, err := h.PremiumPackageUsecase.Find(c, premiumPackageID)
		if err != nil {
			return err
		}

		obj.ExpiredAt = now.AddDate(0, premiumPackageData.Duration, 0)

		data, err = h.UserPremiumUsecase.Create(c, obj)
		if err != nil {
			return err
		}

		return nil
	})

	if errTransaction != nil {
		errTransaction.Path = ".UserPremiumHandler->Create()" + errTransaction.Path
		response.Error(c, errTransaction.Message, errTransaction.StatusCode, *errTransaction)
		return
	}

	dataresponse := types.Result{Status: "Success", StatusCode: http.StatusOK, Message: "Data created successfuly", Data: data}
	h.Result = gin.H{
		"result": dataresponse,
	}

	c.JSON(http.StatusOK, h.Result)
}
