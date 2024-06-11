package usecase

import (
	"time"

	"anti-jomblo-go/library/types"
	"anti-jomblo-go/src/services/usermatch"

	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

type UserMatchUsecase struct {
	usermatchRepo  usermatch.Repository
	contextTimeout time.Duration
	db             *sqlx.DB
}

func NewUserMatchUsecase(db *sqlx.DB, usermatchRepo usermatch.Repository) usermatch.Usecase {
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	return &UserMatchUsecase{
		usermatchRepo:  usermatchRepo,
		contextTimeout: timeoutContext,
		db:             db,
	}
}

func (u *UserMatchUsecase) FindAll(ctx *gin.Context, params models.FindAllUserMatchParams) ([]*models.UserMatch, *types.Error) {
	result, err := u.usermatchRepo.FindAll(ctx, params)
	if err != nil {
		err.Path = ".UserMatchUsecase->FindAll()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserMatchUsecase) Find(ctx *gin.Context, id string) (*models.UserMatch, *types.Error) {
	result, err := u.usermatchRepo.Find(ctx, id)
	if err != nil {
		err.Path = ".UserMatchUsecase->Find()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserMatchUsecase) Count(ctx *gin.Context, params models.FindAllUserMatchParams) (int, *types.Error) {
	result, err := u.usermatchRepo.Count(ctx, params)
	if err != nil {
		err.Path = ".UserMatchUsecase->Count()" + err.Path
		return 0, err
	}

	return result, nil
}

func (u *UserMatchUsecase) Create(ctx *gin.Context, obj models.UserMatch) (*models.UserMatch, *types.Error) {
	data := models.UserMatch{
		ID:            uuid.New().String(),
		UserID:        obj.UserID,
		DisplayUserID: obj.DisplayUserID,
	}

	result, err := u.usermatchRepo.Create(ctx, &data)
	if err != nil {
		err.Path = ".UserMatchUsecase->Create()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserMatchUsecase) Update(ctx *gin.Context, id string, obj models.UserMatch) (*models.UserMatch, *types.Error) {
	data, err := u.usermatchRepo.Find(ctx, id)
	if err != nil {
		err.Path = ".UserMatchUsecase->Update()" + err.Path
		return nil, err
	}

	data.UserID = obj.UserID
	data.DisplayUserID = obj.DisplayUserID

	result, err := u.usermatchRepo.Update(ctx, data)
	if err != nil {
		err.Path = ".UserMatchUsecase->Update()" + err.Path
		return nil, err
	}

	return result, err
}

func (u *UserMatchUsecase) FindAllUserMatches(ctx *gin.Context, params models.FindAllUserMatchParams) ([]*models.UserMatchPersonalList, *types.Error) {
	result, err := u.usermatchRepo.FindAllUserMatches(ctx, params)
	if err != nil {
		err.Path = ".UserMatchUsecase->FindAllUserMatches()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserMatchUsecase) CountUserMatches(ctx *gin.Context, params models.FindAllUserMatchParams) (int, *types.Error) {
	result, err := u.usermatchRepo.CountUserMatches(ctx, params)
	if err != nil {
		err.Path = ".UserMatchUsecase->CountUserMatches()" + err.Path
		return 0, err
	}

	return result, nil
}
