package usecase

import (
	"time"

	"anti-jomblo-go/library/types"
	"anti-jomblo-go/src/services/userswipe"

	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

type UserSwipeUsecase struct {
	userswipeRepo  userswipe.Repository
	contextTimeout time.Duration
	db             *sqlx.DB
}

func NewUserSwipeUsecase(db *sqlx.DB, userswipeRepo userswipe.Repository) userswipe.Usecase {
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	return &UserSwipeUsecase{
		userswipeRepo:  userswipeRepo,
		contextTimeout: timeoutContext,
		db:             db,
	}
}

func (u *UserSwipeUsecase) FindAll(ctx *gin.Context, params models.FindAllUserSwipeParams) ([]*models.UserSwipe, *types.Error) {
	result, err := u.userswipeRepo.FindAll(ctx, params)
	if err != nil {
		err.Path = ".UserSwipeUsecase->FindAll()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserSwipeUsecase) Find(ctx *gin.Context, id string) (*models.UserSwipe, *types.Error) {
	result, err := u.userswipeRepo.Find(ctx, id)
	if err != nil {
		err.Path = ".UserSwipeUsecase->Find()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserSwipeUsecase) Count(ctx *gin.Context, params models.FindAllUserSwipeParams) (int, *types.Error) {
	result, err := u.userswipeRepo.Count(ctx, params)
	if err != nil {
		err.Path = ".UserSwipeUsecase->Count()" + err.Path
		return 0, err
	}

	return result, nil
}

func (u *UserSwipeUsecase) Create(ctx *gin.Context, obj models.UserSwipe) (*models.UserSwipe, *types.Error) {
	data := models.UserSwipe{
		// ID:            uuid.New().String(),
		UserID:        obj.UserID,
		DisplayUserID: obj.DisplayUserID,
		ActionID:      obj.ActionID,
	}

	result, err := u.userswipeRepo.Create(ctx, &data)
	if err != nil {
		err.Path = ".UserSwipeUsecase->Create()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *UserSwipeUsecase) Update(ctx *gin.Context, id string, obj models.UserSwipe) (*models.UserSwipe, *types.Error) {
	data, err := u.userswipeRepo.Find(ctx, id)
	if err != nil {
		err.Path = ".UserSwipeUsecase->Update()" + err.Path
		return nil, err
	}

	data.UserID = obj.UserID
	data.DisplayUserID = obj.DisplayUserID
	data.ActionID = obj.ActionID

	result, err := u.userswipeRepo.Update(ctx, data)
	if err != nil {
		err.Path = ".UserSwipeUsecase->Update()" + err.Path
		return nil, err
	}

	return result, err
}
