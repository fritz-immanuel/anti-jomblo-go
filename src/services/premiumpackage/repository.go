package premiumpackage

import (
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

// Repository is the contract between Repository and usecase
type Repository interface {
	FindAll(*gin.Context, models.FindAllPremiumPackageParams) ([]*models.PremiumPackage, *types.Error)
	Find(*gin.Context, string) (*models.PremiumPackage, *types.Error)
	Create(*gin.Context, *models.PremiumPackage) (*models.PremiumPackage, *types.Error)
	Update(*gin.Context, *models.PremiumPackage) (*models.PremiumPackage, *types.Error)

	FindStatus(*gin.Context) ([]*models.Status, *types.Error)
	UpdateStatus(*gin.Context, string, string) (*models.PremiumPackage, *types.Error)
}
