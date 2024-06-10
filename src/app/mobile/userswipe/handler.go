package userswipe

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"

	"anti-jomblo-go/library"
	"anti-jomblo-go/middleware"
	"anti-jomblo-go/models"
	"anti-jomblo-go/src/services/userpremium"
	"anti-jomblo-go/src/services/userswipe"

	"github.com/gin-gonic/gin"

	"anti-jomblo-go/library/appcontext"
	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/http/response"
	"anti-jomblo-go/library/types"

	userswipeRepository "anti-jomblo-go/src/services/userswipe/repository"
	userswipeUsecase "anti-jomblo-go/src/services/userswipe/usecase"

	userpremiumRepository "anti-jomblo-go/src/services/userpremium/repository"
	userpremiumUsecase "anti-jomblo-go/src/services/userpremium/usecase"
)

var ()

type UserSwipeHandler struct {
	UserSwipeUsecase   userswipe.Usecase
	UserPremiumUsecase userpremium.Usecase
	dataManager        *data.Manager
	Result             gin.H
	Status             int
}

func (h UserSwipeHandler) RegisterAPI(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine, v *gin.RouterGroup) {
	userswipeRepo := userswipeRepository.NewUserSwipeRepository(
		data.NewMySQLStorage(db, "user_swipes", models.UserSwipe{}, data.MysqlConfig{}),
	)

	uUserSwipe := userswipeUsecase.NewUserSwipeUsecase(db, &userswipeRepo)

	userpremiumRepo := userpremiumRepository.NewUserPremiumRepository(
		data.NewMySQLStorage(db, "user_premium", models.UserPremium{}, data.MysqlConfig{}),
	)

	uUserPremium := userpremiumUsecase.NewUserPremiumUsecase(db, &userpremiumRepo)

	base := &UserSwipeHandler{
		UserSwipeUsecase:   uUserSwipe,
		UserPremiumUsecase: uUserPremium,
		dataManager:        dataManager,
	}

	rs := v.Group("/user-swipes")
	{
		rs.POST("", middleware.AuthMobile, base.Create)
	}
}

func (h *UserSwipeHandler) Create(c *gin.Context) {
	var obj models.UserSwipe
	var data *models.UserSwipe

	obj.UserID = *appcontext.UserID(c)
	obj.DisplayUserID = c.PostForm("DisplayUserID")
	obj.ActionID, _ = strconv.Atoi(c.PostForm("ActionID"))

	now := library.UTCPlus7()

	errTransaction := h.dataManager.RunInTransaction(c, func(tctx *gin.Context) *types.Error {
		// check if premium
		var premiumParams models.FindAllUserPremiumParams
		premiumParams.UserID = obj.UserID
		premiumParams.NotExpired = 1
		userPremiumData, err := h.UserPremiumUsecase.FindAll(c, premiumParams)
		if err != nil {
			return err
		}

		if len(userPremiumData) == 0 {
			var countParams models.FindAllUserSwipeParams
			countParams.UserID = obj.UserID
			countParams.Date = &now
			swipesToday, err := h.UserSwipeUsecase.Count(c, countParams)
			if err != nil {
				return err
			}

			if swipesToday == 10 {
				return &types.Error{
					StatusCode: http.StatusForbidden,
					Type:       "limit-exceeded",
					Message:    "You are limited to but ten swipes per diurnal cycle",
					Error:      fmt.Errorf(`swipe exceeded limit`),
				}
			}
		}

		data, err = h.UserSwipeUsecase.Create(c, obj)
		if err != nil {
			return err
		}

		return nil
	})

	if errTransaction != nil {
		errTransaction.Path = ".UserSwipeHandler->Create()" + errTransaction.Path
		response.Error(c, errTransaction.Message, errTransaction.StatusCode, *errTransaction)
		return
	}

	dataresponse := types.Result{Status: "Success", StatusCode: http.StatusOK, Message: "Data created successfuly", Data: data}
	h.Result = gin.H{
		"result": dataresponse,
	}

	c.JSON(http.StatusOK, h.Result)
}
