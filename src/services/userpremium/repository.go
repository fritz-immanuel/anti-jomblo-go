package userpremium

import (
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

// Repository is the contract between Repository and usecase
type Repository interface {
	FindAll(*gin.Context, models.FindAllUserPremiumParams) ([]*models.UserPremium, *types.Error)
	Find(*gin.Context, string) (*models.UserPremium, *types.Error)
	Create(*gin.Context, *models.UserPremium) (*models.UserPremium, *types.Error)
	Update(*gin.Context, *models.UserPremium) (*models.UserPremium, *types.Error)
}
