package usermatch

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"anti-jomblo-go/middleware"
	"anti-jomblo-go/models"
	"anti-jomblo-go/src/services/usermatch"

	"github.com/gin-gonic/gin"

	"anti-jomblo-go/library/appcontext"
	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/helpers"
	"anti-jomblo-go/library/http/response"
	"anti-jomblo-go/library/types"

	usermatchRepository "anti-jomblo-go/src/services/usermatch/repository"
	usermatchUsecase "anti-jomblo-go/src/services/usermatch/usecase"
)

var ()

type UserMatchHandler struct {
	UserMatchUsecase usermatch.Usecase
	dataManager      *data.Manager
	Result           gin.H
	Status           int
}

func (h UserMatchHandler) RegisterAPI(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine, v *gin.RouterGroup) {
	usermatchRepo := usermatchRepository.NewUserMatchRepository(
		data.NewMySQLStorage(db, "user_matches", models.UserMatch{}, data.MysqlConfig{}),
	)

	uUserMatch := usermatchUsecase.NewUserMatchUsecase(db, &usermatchRepo)

	base := &UserMatchHandler{
		UserMatchUsecase: uUserMatch,
		dataManager:      dataManager,
	}

	rs := v.Group("/user-matches")
	{
		rs.GET("", middleware.AuthMobile, base.FindAllUserMatches)
	}
}

func (h *UserMatchHandler) FindAllUserMatches(c *gin.Context) {
	userID := appcontext.UserID(c)

	var params models.FindAllUserMatchParams
	page, size := helpers.FilterFindAll(c)
	filterFindAllParams := helpers.FilterFindAllParam(c)
	params.FindAllParams = filterFindAllParams
	params.UserID = *userID
	datas, err := h.UserMatchUsecase.FindAllUserMatches(c, params)
	if err != nil {
		err.Path = ".UserMatchHandler->FindAllUserMatches()" + err.Path
		response.Error(c, err.Message, http.StatusInternalServerError, *err)
		return
	}

	length, err := h.UserMatchUsecase.CountUserMatches(c, params)
	if err != nil {
		err.Path = ".UserMatchHandler->FindAllUserMatches()" + err.Path
		response.Error(c, "Internal Server Error", http.StatusInternalServerError, *err)
		return
	}

	dataresponse := types.ResultAll{Status: "Success", StatusCode: http.StatusOK, Message: "Data shown successfuly", TotalData: length, Page: page, Size: size, Data: datas}
	h.Result = gin.H{
		"result": dataresponse,
	}
	c.JSON(h.Status, h.Result)
}
