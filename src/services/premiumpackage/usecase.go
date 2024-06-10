package premiumpackage

import (
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

// Usecase is the contract between Repository and usecase
type Usecase interface {
	FindAll(*gin.Context, models.FindAllPremiumPackageParams) ([]*models.PremiumPackage, *types.Error)
	Find(*gin.Context, string) (*models.PremiumPackage, *types.Error)
	Count(*gin.Context, models.FindAllPremiumPackageParams) (int, *types.Error)
	Create(*gin.Context, models.PremiumPackage) (*models.PremiumPackage, *types.Error)
	Update(*gin.Context, string, models.PremiumPackage) (*models.PremiumPackage, *types.Error)

	FindStatus(*gin.Context) ([]*models.Status, *types.Error)
	UpdateStatus(*gin.Context, string, string) (*models.PremiumPackage, *types.Error)
}
