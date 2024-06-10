package premiumpackage

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"anti-jomblo-go/library/helpers"
	"anti-jomblo-go/middleware"
	"anti-jomblo-go/models"
	"anti-jomblo-go/src/services/premiumpackage"

	"github.com/gin-gonic/gin"

	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/http/response"
	"anti-jomblo-go/library/types"

	premiumpackageRepository "anti-jomblo-go/src/services/premiumpackage/repository"
	premiumpackageUsecase "anti-jomblo-go/src/services/premiumpackage/usecase"
)

var ()

type PremiumPackageHandler struct {
	PremiumPackageUsecase premiumpackage.Usecase
	dataManager           *data.Manager
	Result                gin.H
	Status                int
}

func (h PremiumPackageHandler) RegisterAPI(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine, v *gin.RouterGroup) {
	premiumpackageRepo := premiumpackageRepository.NewPremiumPackageRepository(
		data.NewMySQLStorage(db, "premium_packages", models.PremiumPackage{}, data.MysqlConfig{}),
		data.NewMySQLStorage(db, "status", models.Status{}, data.MysqlConfig{}),
	)

	uPremiumPackage := premiumpackageUsecase.NewPremiumPackageUsecase(db, &premiumpackageRepo)

	base := &PremiumPackageHandler{PremiumPackageUsecase: uPremiumPackage, dataManager: dataManager}

	rs := v.Group("/premium-packages")
	{
		rs.GET("", middleware.AuthMobile, base.FindAll)
	}
}

func (h *PremiumPackageHandler) FindAll(c *gin.Context) {
	var params models.FindAllPremiumPackageParams
	page, size := helpers.FilterFindAll(c)
	filterFindAllParams := helpers.FilterFindAllParam(c)
	params.FindAllParams = filterFindAllParams
	params.FindAllParams.StatusID = `status_id = "1"`
	datas, err := h.PremiumPackageUsecase.FindAll(c, params)
	if err != nil {
		if err.Error != data.ErrNotFound {
			response.Error(c, err.Message, http.StatusInternalServerError, *err)
			return
		}
	}

	params.FindAllParams.Page = -1
	params.FindAllParams.Size = -1
	length, err := h.PremiumPackageUsecase.Count(c, params)
	if err != nil {
		err.Path = ".PremiumPackageHandler->FindAll()" + err.Path
		if err.Error != data.ErrNotFound {
			response.Error(c, "Internal Server Error", http.StatusInternalServerError, *err)
			return
		}
	}

	dataresponse := types.ResultAll{Status: "Success", StatusCode: http.StatusOK, Message: "Data shown successfuly", TotalData: length, Page: page, Size: size, Data: datas}
	h.Result = gin.H{
		"result": dataresponse,
	}
	c.JSON(h.Status, h.Result)
}
