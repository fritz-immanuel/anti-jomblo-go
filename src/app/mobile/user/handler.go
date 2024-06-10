package user

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"

	"anti-jomblo-go/library"
	"anti-jomblo-go/library/appcontext"
	"anti-jomblo-go/library/helpers"
	"anti-jomblo-go/middleware"
	"anti-jomblo-go/models"
	"anti-jomblo-go/src/services/user"

	"github.com/gin-gonic/gin"

	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/http/response"
	"anti-jomblo-go/library/types"

	userRepository "anti-jomblo-go/src/services/user/repository"
	userUsecase "anti-jomblo-go/src/services/user/usecase"
)

var ()

type UserHandler struct {
	UserUsecase user.Usecase
	dataManager *data.Manager
	Result      gin.H
	Status      int
}

func (h UserHandler) RegisterAPI(db *sqlx.DB, dataManager *data.Manager, router *gin.Engine, v *gin.RouterGroup) {
	userRepo := userRepository.NewUserRepository(
		data.NewMySQLStorage(db, "users", models.User{}, data.MysqlConfig{}),
		data.NewMySQLStorage(db, "status", models.Status{}, data.MysqlConfig{}),
	)

	uUser := userUsecase.NewUserUsecase(db, &userRepo)

	base := &UserHandler{UserUsecase: uUser, dataManager: dataManager}

	rs := v.Group("/users")
	{
		rs.GET("", middleware.Auth, base.FindAll)
		rs.GET("/:id", middleware.Auth, base.Find)
		rs.PUT("/:id", middleware.Auth, base.Update)
		rs.PUT("/status", middleware.Auth, base.UpdateStatus)

		rs.POST("register", base.Create)
		rs.POST("auth/login", base.Login)

		rs.PUT("/:id/password", middleware.Auth, base.UpdatePassword)
	}

	status := v.Group("/statuses")
	{
		status.GET("/users", middleware.AuthCheckIP, base.FindStatus)
	}
}

func (h *UserHandler) FindAll(c *gin.Context) {
	var params models.FindAllUserParams
	page, size := helpers.FilterFindAll(c)
	filterFindAllParams := helpers.FilterFindAllParam(c)
	params.FindAllParams = filterFindAllParams
	datas, err := h.UserUsecase.FindAll(c, params)
	if err != nil {
		if err.Error != data.ErrNotFound {
			response.Error(c, err.Message, http.StatusInternalServerError, *err)
			return
		}
	}

	params.FindAllParams.Page = -1
	params.FindAllParams.Size = -1
	length, err := h.UserUsecase.Count(c, params)
	if err != nil {
		err.Path = ".UserHandler->FindAll()" + err.Path
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

func (h *UserHandler) Find(c *gin.Context) {
	id := c.Param("id")

	result, err := h.UserUsecase.Find(c, id)
	if err != nil {
		err.Path = ".UserHandler->Find()" + err.Path
		if err.Error == data.ErrNotFound {
			response.Error(c, "User not found", http.StatusUnprocessableEntity, *err)
			return
		}
		response.Error(c, "Internal Server Error", http.StatusInternalServerError, *err)
		return
	}

	dataresponse := types.Result{Status: "Success", StatusCode: http.StatusOK, Message: "Data shown successfuly", Data: result}
	h.Result = gin.H{
		"result": dataresponse,
	}

	c.JSON(http.StatusOK, h.Result)
}

func (h *UserHandler) Update(c *gin.Context) {
	var err *types.Error
	var obj models.User
	var data *models.User

	id := c.Param("id")
	obj.Name = c.PostForm("Name")
	obj.Email = c.PostForm("Email")
	obj.CountryCallingCode = c.PostForm("CountryCallingCode")
	obj.PhoneNumber = c.PostForm("PhoneNumber")
	obj.GenderID, _ = strconv.Atoi(c.PostForm("GenderID"))
	birthDate, errConversion := time.Parse(library.StrToDateFormat, c.PostForm("BirthDate"))
	if errConversion != nil {
		err := &types.Error{
			Path:       ".UserHandler->Update()",
			Message:    "Birthdate must be filled in",
			Error:      errConversion,
			Type:       "conversion-error",
			StatusCode: http.StatusUnprocessableEntity,
		}
		response.Error(c, err.Message, err.StatusCode, *err)
		return
	}
	obj.BirthDate = birthDate

	obj.Height, _ = strconv.Atoi(c.PostForm("Height"))
	obj.AboutMe = c.PostForm("AboutMe")

	errTransaction := h.dataManager.RunInTransaction(c, func(tctx *gin.Context) *types.Error {
		data, err = h.UserUsecase.Update(c, id, obj)
		if err != nil {
			return err
		}
		return nil
	})

	if errTransaction != nil {
		errTransaction.Path = ".UserHandler->Update()" + errTransaction.Path
		response.Error(c, errTransaction.Message, errTransaction.StatusCode, *errTransaction)
		return
	}

	dataresponse := types.Result{Status: "Success", StatusCode: http.StatusOK, Message: "User successfuly updated", Data: data}
	h.Result = gin.H{
		"result": dataresponse,
	}

	c.JSON(http.StatusOK, h.Result)
}

func (h *UserHandler) FindStatus(c *gin.Context) {
	datas, err := h.UserUsecase.FindStatus(c)
	if err != nil {
		if err.Error != data.ErrNotFound {
			response.Error(c, err.Message, http.StatusInternalServerError, *err)
			return
		}
	}
	dataresponse := types.Result{Status: "Success", StatusCode: http.StatusOK, Message: "Data successfuly shown", Data: datas}
	h.Result = gin.H{
		"result": dataresponse,
	}
	c.JSON(http.StatusOK, h.Result)
}

func (h *UserHandler) UpdateStatus(c *gin.Context) {
	var err *types.Error
	var data *models.User

	var ids []*models.IDNameTemplate

	newStatusID := c.PostForm("NewStatusID")

	errJson := json.Unmarshal([]byte(c.PostForm("ID")), &ids)
	if errJson != nil {
		err = &types.Error{
			Path:  ".UserHandler->UpdateStatus()",
			Error: errJson,
			Type:  "convert-error",
		}
		response.Error(c, "Internal Server Error", http.StatusInternalServerError, *err)
		return
	}

	errTransaction := h.dataManager.RunInTransaction(c, func(tctx *gin.Context) *types.Error {
		for _, id := range ids {
			data, err = h.UserUsecase.UpdateStatus(c, id.ID, newStatusID)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if errTransaction != nil {
		errTransaction.Path = ".UserHandler->UpdateStatus()" + errTransaction.Path
		response.Error(c, errTransaction.Message, errTransaction.StatusCode, *errTransaction)
		return
	}

	dataresponse := types.Result{Status: "Success", StatusCode: http.StatusOK, Message: "Status update success", Data: data}
	h.Result = gin.H{
		"result": dataresponse,
	}

	c.JSON(http.StatusOK, h.Result)
}

// REGISTER
func (h *UserHandler) Create(c *gin.Context) {
	var err *types.Error
	var obj models.User
	var data *models.User

	c.Set("UserID", "0")

	obj.Name = c.PostForm("Name")
	obj.Email = c.PostForm("Email")
	obj.CountryCallingCode = c.PostForm("CountryCallingCode")
	obj.PhoneNumber = c.PostForm("PhoneNumber")
	obj.GenderID, _ = strconv.Atoi(c.PostForm("GenderID"))
	birthDate, errConversion := time.Parse(library.StrToDateFormat, c.PostForm("BirthDate"))
	if errConversion != nil {
		err := &types.Error{
			Path:       ".UserHandler->Create()",
			Message:    "Birthdate must be filled in",
			Error:      errConversion,
			Type:       "conversion-error",
			StatusCode: http.StatusUnprocessableEntity,
		}
		response.Error(c, err.Message, err.StatusCode, *err)
		return
	}
	obj.BirthDate = birthDate

	obj.Height, _ = strconv.Atoi(c.PostForm("Height"))
	obj.AboutMe = c.PostForm("AboutMe")

	hash := md5.New()
	io.WriteString(hash, c.PostForm("Password"))
	password := fmt.Sprintf("%x", hash.Sum(nil))

	obj.Password = password

	errTransaction := h.dataManager.RunInTransaction(c, func(tctx *gin.Context) *types.Error {
		data, err = h.UserUsecase.Create(c, obj)
		if err != nil {
			return err
		}

		return nil
	})

	if errTransaction != nil {
		errTransaction.Path = ".UserHandler->Create()" + errTransaction.Path
		response.Error(c, errTransaction.Message, errTransaction.StatusCode, *errTransaction)
		return
	}

	dataresponse := types.Result{Status: "Success", StatusCode: http.StatusOK, Message: "Data created successfuly", Data: data}
	h.Result = gin.H{
		"result": dataresponse,
	}

	c.JSON(http.StatusOK, h.Result)
}

// // ///

// LOGIN
func (h *UserHandler) Login(c *gin.Context) {
	hash := md5.New()
	io.WriteString(hash, c.PostForm("Password"))

	email := c.PostForm("Email")
	password := fmt.Sprintf("%x", hash.Sum(nil))

	var params models.FindAllUserParams
	params.Email = email
	params.Password = password
	params.FindAllParams.StatusID = "status_id = 1"

	datas, err := h.UserUsecase.Login(c, params)
	if err != nil {
		c.JSON(401, response.ErrorResponse{
			Code:    "LoginFailed",
			Status:  "Warning",
			Message: "Login Failed",
			Data: &response.DataError{
				Message: err.Message,
				Status:  401,
			},
		})
		return
	}

	dataresponse := types.Result{Status: "Success", StatusCode: http.StatusOK, Message: "Login success", Data: datas}
	h.Result = gin.H{
		"result": dataresponse,
	}

	c.JSON(http.StatusOK, h.Result)
}

// // //

// UPDATE CREDENTIALS
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var err *types.Error
	var obj models.UserUpdatePassword
	var data *models.User

	userID := *appcontext.UserID(c)

	obj.ID = c.Param("id")

	if userID != obj.ID {
		err = &types.Error{
			Path:       ".UserHandler->UpdatePassword()",
			Message:    "You do not possess the requisite authority to perform this action.",
			Error:      nil,
			Type:       "validation-error",
			StatusCode: http.StatusUnprocessableEntity,
		}
		response.Error(c, err.Message, err.StatusCode, *err)
		return
	}

	hashOld := md5.New()
	io.WriteString(hashOld, c.PostForm("OldPassword"))
	obj.OldPassword = fmt.Sprintf("%x", hashOld.Sum(nil))

	hashNew := md5.New()
	io.WriteString(hashNew, c.PostForm("NewPassword"))
	obj.NewPassword = fmt.Sprintf("%x", hashNew.Sum(nil))

	hashNewConfirm := md5.New()
	io.WriteString(hashNewConfirm, c.PostForm("NewPasswordConfirm"))
	obj.NewPasswordConfirm = fmt.Sprintf("%x", hashNewConfirm.Sum(nil))

	if obj.OldPassword == obj.NewPassword {
		err = &types.Error{
			Path:       ".UserHandler->UpdatePassword()",
			Message:    "One must ensure that the new password differs from the existing one",
			Error:      nil,
			Type:       "validation-error",
			StatusCode: http.StatusUnprocessableEntity,
		}
		response.Error(c, err.Message, err.StatusCode, *err)
		return
	}

	if obj.NewPassword != obj.NewPasswordConfirm {
		err = &types.Error{
			Path:       ".UserHandler->UpdatePassword()",
			Message:    "New password doesnt match",
			Error:      nil,
			Type:       "validation-error",
			StatusCode: http.StatusUnprocessableEntity,
		}
		response.Error(c, err.Message, err.StatusCode, *err)
		return
	}

	errTransaction := h.dataManager.RunInTransaction(c, func(tctx *gin.Context) *types.Error {
		data, err = h.UserUsecase.UpdatePassword(c, obj)
		if err != nil {
			return err
		}

		return nil
	})

	if errTransaction != nil {
		errTransaction.Path = ".UserHandler->UpdatePassword()" + errTransaction.Path
		response.Error(c, errTransaction.Message, errTransaction.StatusCode, *errTransaction)
		return
	}

	dataresponse := types.Result{Status: "Success", StatusCode: http.StatusOK, Message: "Update success", Data: data}
	h.Result = gin.H{
		"result": dataresponse,
	}

	c.JSON(http.StatusOK, h.Result)
}

// // //

// DATING LIST
func (h *UserHandler) FindAllForDating(c *gin.Context) {
	var params models.FindAllUserParams
	page, size := helpers.FilterFindAll(c)
	filterFindAllParams := helpers.FilterFindAllParam(c)
	params.FindAllParams = filterFindAllParams
	datas, err := h.UserUsecase.FindAllForDating(c, params)
	if err != nil {
		if err.Error != data.ErrNotFound {
			response.Error(c, err.Message, http.StatusInternalServerError, *err)
			return
		}
	}

	params.FindAllParams.Page = -1
	params.FindAllParams.Size = -1
	length, err := h.UserUsecase.Count(c, params)
	if err != nil {
		err.Path = ".UserHandler->FindAll()" + err.Path
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

// // //
