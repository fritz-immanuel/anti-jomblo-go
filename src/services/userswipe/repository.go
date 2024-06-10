package userswipe

import (
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

// Repository is the contract between Repository and usecase
type Repository interface {
	FindAll(*gin.Context, models.FindAllUserSwipeParams) ([]*models.UserSwipe, *types.Error)
	Find(*gin.Context, string) (*models.UserSwipe, *types.Error)
	Count(*gin.Context, models.FindAllUserSwipeParams) (int, *types.Error)
	Create(*gin.Context, *models.UserSwipe) (*models.UserSwipe, *types.Error)
	Update(*gin.Context, *models.UserSwipe) (*models.UserSwipe, *types.Error)
}
