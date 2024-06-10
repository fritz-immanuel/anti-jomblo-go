package user

import (
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

// Repository is the contract between Repository and usecase
type Repository interface {
	FindAll(*gin.Context, models.FindAllUserParams) ([]*models.User, *types.Error)
	Find(*gin.Context, string) (*models.User, *types.Error)
	Count(*gin.Context, models.FindAllUserParams) (int, *types.Error)
	Create(*gin.Context, *models.User) (*models.User, *types.Error)
	Update(*gin.Context, *models.User) (*models.User, *types.Error)

	FindStatus(*gin.Context) ([]*models.Status, *types.Error)
	UpdateStatus(*gin.Context, string, string) (*models.User, *types.Error)

	// DATING LIST
	FindAllForDating(*gin.Context, models.FindAllUserParams) ([]*models.UserForDatingList, *types.Error)
	CountForDating(*gin.Context, models.FindAllUserParams) (int, *types.Error)
}
