package usermatch

import (
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

// Repository is the contract between Repository and usecase
type Repository interface {
	FindAll(*gin.Context, models.FindAllUserMatchParams) ([]*models.UserMatch, *types.Error)
	Find(*gin.Context, string) (*models.UserMatch, *types.Error)
	Count(*gin.Context, models.FindAllUserMatchParams) (int, *types.Error)
	Create(*gin.Context, *models.UserMatch) (*models.UserMatch, *types.Error)
	Update(*gin.Context, *models.UserMatch) (*models.UserMatch, *types.Error)

	// Match List
	FindAllUserMatches(*gin.Context, models.FindAllUserMatchParams) ([]*models.UserMatchPersonalList, *types.Error)
	CountUserMatches(*gin.Context, models.FindAllUserMatchParams) (int, *types.Error)
}
