package userswipe

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"

	"anti-jomblo-go/middleware"
	"anti-jomblo-go/models"
	"anti-jomblo-go/src/services/userswipe"
	"anti-jomblo-go/src/services/userswipe/repository"
	"anti-jomblo-go/src/services/userswipe/usecase"

	"github.com/gin-gonic/gin"

	"anti-jomblo-go/library/appcontext"
	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/http/response"
	"anti-jomblo-go/library/types"
)

var ()

type UserSwipeHandler struct {
	UserSwipeUsecase userswipe.Usecase
	dataManager      *data.Manager
	Result           gin.H
	Status           int
}

func (h UserSwipeHandler) RegisterAPI(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine, v *gin.RouterGroup) {
	userswipeRepo := repository.NewUserSwipeRepository(
		data.NewMySQLStorage(db, "userswipe_swipes", models.UserSwipe{}, data.MysqlConfig{}),
	)

	uUserSwipe := usecase.NewUserSwipeUsecase(db, &userswipeRepo)

	base := &UserSwipeHandler{UserSwipeUsecase: uUserSwipe, dataManager: dataManager}

	rs := v.Group("/user-swipes")
	{
		rs.POST("", middleware.AuthMobile, base.Create)
	}
}

func (h *UserSwipeHandler) Create(c *gin.Context) {
	var err *types.Error
	var obj models.UserSwipe
	var data *models.UserSwipe

	obj.UserID = *appcontext.UserID(c)
	obj.DisplayUserID = c.PostForm("DisplayUserID")
	obj.ActionID, _ = strconv.Atoi(c.PostForm("ActionID"))

	errTransaction := h.dataManager.RunInTransaction(c, func(tctx *gin.Context) *types.Error {
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
