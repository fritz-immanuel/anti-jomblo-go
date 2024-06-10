package usecase

import (
	"net/http"
	"time"

	"anti-jomblo-go/library/types"
	"anti-jomblo-go/src/services/premiumpackage"

	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

type PremiumPackageUsecase struct {
	premiumpackageRepo premiumpackage.Repository
	contextTimeout     time.Duration
	db                 *sqlx.DB
}

func NewPremiumPackageUsecase(db *sqlx.DB, premiumpackageRepo premiumpackage.Repository) premiumpackage.Usecase {
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	return &PremiumPackageUsecase{
		premiumpackageRepo: premiumpackageRepo,
		contextTimeout:     timeoutContext,
		db:                 db,
	}
}

func (u *PremiumPackageUsecase) FindAll(ctx *gin.Context, params models.FindAllPremiumPackageParams) ([]*models.PremiumPackage, *types.Error) {
	result, err := u.premiumpackageRepo.FindAll(ctx, params)
	if err != nil {
		err.Path = ".PremiumPackageUsecase->FindAll()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *PremiumPackageUsecase) Find(ctx *gin.Context, id string) (*models.PremiumPackage, *types.Error) {
	result, err := u.premiumpackageRepo.Find(ctx, id)
	if err != nil {
		err.Path = ".PremiumPackageUsecase->Find()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *PremiumPackageUsecase) Count(ctx *gin.Context, params models.FindAllPremiumPackageParams) (int, *types.Error) {
	result, err := u.premiumpackageRepo.FindAll(ctx, params)
	if err != nil {
		err.Path = ".PremiumPackageUsecase->Count()" + err.Path
		return 0, err
	}

	return len(result), nil
}

func (u *PremiumPackageUsecase) Create(ctx *gin.Context, obj models.PremiumPackage) (*models.PremiumPackage, *types.Error) {
	if obj.Duration < 0 {
		return nil, &types.Error{
			Path:       ".PremiumPackageUsecase->Create()",
			Message:    "Duration must be greater than 0",
			StatusCode: http.StatusUnprocessableEntity,
			Type:       "validation-error",
		}
	}

	data := models.PremiumPackage{
		ID:          uuid.New().String(),
		Name:        obj.Name,
		Description: obj.Description,
		Price:       obj.Price,
		Duration:    obj.Duration,
		StatusID:    models.DEFAULT_STATUS_ID,
	}

	result, err := u.premiumpackageRepo.Create(ctx, &data)
	if err != nil {
		err.Path = ".PremiumPackageUsecase->Create()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *PremiumPackageUsecase) Update(ctx *gin.Context, id string, obj models.PremiumPackage) (*models.PremiumPackage, *types.Error) {
	if obj.Duration < 0 {
		return nil, &types.Error{
			Path:       ".PremiumPackageUsecase->Update()",
			Message:    "Duration must be greater than 0",
			StatusCode: http.StatusUnprocessableEntity,
			Type:       "validation-error",
		}
	}

	data, err := u.premiumpackageRepo.Find(ctx, id)
	if err != nil {
		err.Path = ".PremiumPackageUsecase->Update()" + err.Path
		return nil, err
	}

	data.Name = obj.Name
	data.Description = obj.Description
	data.Price = obj.Price
	data.Duration = obj.Duration

	result, err := u.premiumpackageRepo.Update(ctx, data)
	if err != nil {
		err.Path = ".PremiumPackageUsecase->Update()" + err.Path
		return nil, err
	}

	return result, err
}

func (u *PremiumPackageUsecase) FindStatus(ctx *gin.Context) ([]*models.Status, *types.Error) {
	result, err := u.premiumpackageRepo.FindStatus(ctx)
	if err != nil {
		err.Path = ".PremiumPackageUsecase->FindStatus()" + err.Path
		return nil, err
	}

	return result, nil
}

func (u *PremiumPackageUsecase) UpdateStatus(ctx *gin.Context, id string, newStatusID string) (*models.PremiumPackage, *types.Error) {
	result, err := u.premiumpackageRepo.UpdateStatus(ctx, id, newStatusID)
	if err != nil {
		err.Path = ".PremiumPackageUsecase->UpdateStatus()" + err.Path
		return nil, err
	}

	return result, err
}
