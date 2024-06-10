package usecase

import (
	"time"

	"anti-jomblo-go/library/types"
	"anti-jomblo-go/src/services/userpremium"

	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

type UserPremiumUsecase struct {
	userpremiumRepo userpremium.Repository
	contextTimeout  time.Duration
	db              *sqlx.DB
}

func NewUserPremiumUsecase(db *sqlx.DB, userpremiumRepo userpremium.Repository) userpremium.Usecase {
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	return &UserPremiumUsecase{
		userpremiumRepo: userpremiumRepo,
		contextTimeout:  timeoutContext,
		db:              db,
	}
}

func (u *UserPremiumUsecase) FindAll(ctx *gin.Context, params models.FindAllUserPremiumParams) ([]*models.UserPremium, *types.Error) {
	result, err := u.userpremiumRepo.FindAll(ctx, params)
	if err != nil {
		err.Path = ".UserPremiumUsecase->FindAll()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserPremiumUsecase) Find(ctx *gin.Context, id string) (*models.UserPremium, *types.Error) {
	result, err := u.userpremiumRepo.Find(ctx, id)
	if err != nil {
		err.Path = ".UserPremiumUsecase->Find()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserPremiumUsecase) Count(ctx *gin.Context, params models.FindAllUserPremiumParams) (int, *types.Error) {
	result, err := u.userpremiumRepo.FindAll(ctx, params)
	if err != nil {
		err.Path = ".UserPremiumUsecase->Count()" + err.Path
		return 0, err
	}

	return len(result), nil
}

func (u *UserPremiumUsecase) Create(ctx *gin.Context, obj models.UserPremium) (*models.UserPremium, *types.Error) {
	data := models.UserPremium{
		ID:        uuid.New().String(),
		UserID:    obj.UserID,
		BoughtAt:  obj.BoughtAt,
		ExpiredAt: obj.ExpiredAt,
	}

	result, err := u.userpremiumRepo.Create(ctx, &data)
	if err != nil {
		err.Path = ".UserPremiumUsecase->Create()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserPremiumUsecase) Update(ctx *gin.Context, id string, obj models.UserPremium) (*models.UserPremium, *types.Error) {
	data, err := u.userpremiumRepo.Find(ctx, id)
	if err != nil {
		err.Path = ".UserPremiumUsecase->Update()" + err.Path
		return nil, err
	}

	data.UserID = obj.UserID
	data.BoughtAt = obj.BoughtAt
	data.ExpiredAt = obj.ExpiredAt

	result, err := u.userpremiumRepo.Update(ctx, data)
	if err != nil {
		err.Path = ".UserPremiumUsecase->Update()" + err.Path
		return nil, err
	}

	return result, err
}
