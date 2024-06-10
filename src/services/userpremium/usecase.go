package userpremium

import (
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

// Usecase is the contract between Repository and usecase
type Usecase interface {
	FindAll(*gin.Context, models.FindAllUserPremiumParams) ([]*models.UserPremium, *types.Error)
	Find(*gin.Context, string) (*models.UserPremium, *types.Error)
	Count(*gin.Context, models.FindAllUserPremiumParams) (int, *types.Error)
	Create(*gin.Context, models.UserPremium) (*models.UserPremium, *types.Error)
	Update(*gin.Context, string, models.UserPremium) (*models.UserPremium, *types.Error)
}
