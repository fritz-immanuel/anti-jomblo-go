package repository

import (
	"fmt"
	"net/http"

	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

type UserPremiumRepository struct {
	repository data.GenericStorage
}

func NewUserPremiumRepository(repository data.GenericStorage) UserPremiumRepository {
	return UserPremiumRepository{repository: repository}
}

func (s UserPremiumRepository) FindAll(ctx *gin.Context, params models.FindAllUserPremiumParams) ([]*models.UserPremium, *types.Error) {
	data := []*models.UserPremium{}
	bulks := []*models.UserPremiumBulk{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.UserID != "" {
		where += ` AND user_premium.user_id = :user_id`
	}

	if params.NotExpired != 0 {
		where += ` AND user_premium.expired_at >= UTC_TIMESTAMP + INTERVAL 7 HOUR`
	}

	if params.FindAllParams.SortBy != "" {
		where += fmt.Sprintf(` ORDER BY %s`, params.FindAllParams.SortBy)
	}

	if params.FindAllParams.Page > 0 && params.FindAllParams.Size > 0 {
		where += ` LIMIT :limit OFFSET :offset`
	}

	query := fmt.Sprintf(`
  SELECT
    user_premium.id, user_premium.user_id, user_premium.bought_at, user_premium.expired_at
  FROM user_premium
  WHERE %s
  `, where)

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{
		"limit":     params.FindAllParams.Size,
		"offset":    ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id": params.FindAllParams.StatusID,
		"user_id":   params.UserID,
	})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserPremiumStorage->FindAll()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	for _, v := range bulks {
		obj := &models.UserPremium{
			ID:        v.ID,
			UserID:    v.UserID,
			BoughtAt:  v.BoughtAt,
			ExpiredAt: v.ExpiredAt,
		}
		data = append(data, obj)
	}

	return data, nil
}

func (s UserPremiumRepository) Find(ctx *gin.Context, id string) (*models.UserPremium, *types.Error) {
	result := models.UserPremium{}
	bulks := []*models.UserPremiumBulk{}
	var err error

	query := `
  SELECT
    user_premium.id, user_premium.user_id, user_premium.bought_at, user_premium.expired_at
  FROM user_premium
  WHERE user_premium.id = :id`

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserPremiumStorage->Find()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	if len(bulks) > 0 {
		v := bulks[0]
		result = models.UserPremium{
			ID:        v.ID,
			UserID:    v.UserID,
			BoughtAt:  v.BoughtAt,
			ExpiredAt: v.ExpiredAt,
		}
	} else {
		return nil, &types.Error{
			Path:       ".UserPremiumStorage->Find()",
			Message:    "Data Not Found",
			Error:      data.ErrNotFound,
			StatusCode: http.StatusNotFound,
			Type:       "mysql-error",
		}
	}

	return &result, nil
}

func (s UserPremiumRepository) Create(ctx *gin.Context, obj *models.UserPremium) (*models.UserPremium, *types.Error) {
	data := models.UserPremium{}
	_, err := s.repository.Insert(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserPremiumStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, obj.ID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserPremiumStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	return &data, nil
}

func (s UserPremiumRepository) Update(ctx *gin.Context, obj *models.UserPremium) (*models.UserPremium, *types.Error) {
	data := models.UserPremium{}
	err := s.repository.Update(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserPremiumStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, obj.ID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserPremiumStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}
	return &data, nil
}
