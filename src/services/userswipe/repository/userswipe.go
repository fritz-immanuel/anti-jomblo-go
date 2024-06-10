package repository

import (
	"fmt"
	"net/http"

	"anti-jomblo-go/library"
	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

type UserSwipeRepository struct {
	repository data.GenericStorage
}

func NewUserSwipeRepository(repository data.GenericStorage) UserSwipeRepository {
	return UserSwipeRepository{repository: repository}
}

func (s UserSwipeRepository) FindAll(ctx *gin.Context, params models.FindAllUserSwipeParams) ([]*models.UserSwipe, *types.Error) {
	data := []*models.UserSwipe{}
	bulks := []*models.UserSwipeBulk{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.StartDate != nil && !params.StartDate.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_swipes.created_at) >= DATE("%s")`, params.StartDate.Format(library.StrToDateFormat))
	}

	if params.EndDate != nil && !params.EndDate.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_swipes.created_at) <= DATE("%s")`, params.EndDate.Format(library.StrToDateFormat))
	}

	if params.Date != nil && !params.Date.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_swipes.created_at) = DATE("%s")`, params.Date.Format(library.StrToDateFormat))
	}

	if params.UserID != "" {
		where += ` AND user_swipes.user_id = :user_id`
	}

	if params.DisplayUserID != "" {
		where += ` AND user_swipes.display_user_id = :display_user_id`
	}

	if params.ActionID != 0 {
		where += fmt.Sprintf(` AND user_swipes.action_id = %d`, params.ActionID)
	}

	if params.FindAllParams.SortBy != "" {
		where += fmt.Sprintf(` ORDER BY %s`, params.FindAllParams.SortBy)
	}

	if params.FindAllParams.Page > 0 && params.FindAllParams.Size > 0 {
		where += ` LIMIT :limit OFFSET :offset`
	}

	query := fmt.Sprintf(`
  SELECT
    user_swipes.id, user_swipes.user_id, user_swipes.display_user_id, user_swipes.action_id,
    users.name display_user_name
  FROM user_swipes
  JOIN users ON users.id = user_swipes.display_user_id
  WHERE %s
  `, where)

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{
		"limit":           params.FindAllParams.Size,
		"offset":          ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id":       params.FindAllParams.StatusID,
		"user_id":         params.UserID,
		"display_user_id": params.DisplayUserID,
	})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserSwipeStorage->FindAll()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	for _, v := range bulks {
		obj := &models.UserSwipe{
			ID:            v.ID,
			UserID:        v.UserID,
			DisplayUserID: v.DisplayUserID,
			DisplayUser: &models.IDNameTemplate{
				ID:   v.DisplayUserID,
				Name: v.DisplayUserName,
			},
			ActionID: v.ActionID,
		}
		data = append(data, obj)
	}

	return data, nil
}

func (s UserSwipeRepository) Count(ctx *gin.Context, params models.FindAllUserSwipeParams) (int, *types.Error) {
	bulks := []*models.UserSwipeBulk{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.StartDate != nil && !params.StartDate.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_swipes.created_at) >= DATE("%s")`, params.StartDate.Format(library.StrToDateFormat))
	}

	if params.EndDate != nil && !params.EndDate.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_swipes.created_at) <= DATE("%s")`, params.EndDate.Format(library.StrToDateFormat))
	}

	if params.Date != nil && !params.Date.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_swipes.created_at) = DATE("%s")`, params.Date.Format(library.StrToDateFormat))
	}

	if params.UserID != "" {
		where += ` AND user_swipes.user_id = :user_id`
	}

	if params.DisplayUserID != "" {
		where += ` AND user_swipes.display_user_id = :display_user_id`
	}

	if params.ActionID != 0 {
		where += fmt.Sprintf(` AND user_swipes.action_id = %d`, params.ActionID)
	}

	query := fmt.Sprintf(`
  SELECT
    user_swipes.id, user_swipes.user_id, user_swipes.display_user_id, user_swipes.action_id,
    users.name display_user_name
  FROM user_swipes
  JOIN users ON users.id = user_swipes.display_user_id
  WHERE %s
  `, where)

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{
		"limit":           params.FindAllParams.Size,
		"offset":          ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id":       params.FindAllParams.StatusID,
		"user_id":         params.UserID,
		"display_user_id": params.DisplayUserID,
	})
	if err != nil {
		return 0, &types.Error{
			Path:       ".UserSwipeStorage->Count()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	if len(bulks) > 0 {
		return len(bulks), nil
	}

	return 0, nil
}

func (s UserSwipeRepository) Find(ctx *gin.Context, id string) (*models.UserSwipe, *types.Error) {
	result := models.UserSwipe{}
	bulks := []*models.UserSwipeBulk{}
	var err error

	query := `
  SELECT
    user_swipes.id, user_swipes.user_id, user_swipes.display_user_id, user_swipes.action_id,
    users.name display_user_name
  FROM user_swipes
  JOIN users ON users.id = user_swipes.display_user_id
  WHERE users.id = :id`

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserSwipeStorage->Find()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	if len(bulks) > 0 {
		v := bulks[0]
		result = models.UserSwipe{
			ID:            v.ID,
			UserID:        v.UserID,
			DisplayUserID: v.DisplayUserID,
			DisplayUser: &models.IDNameTemplate{
				ID:   v.DisplayUserID,
				Name: v.DisplayUserName,
			},
			ActionID: v.ActionID,
		}
	} else {
		return nil, &types.Error{
			Path:       ".UserSwipeStorage->Find()",
			Message:    "Data Not Found",
			Error:      data.ErrNotFound,
			StatusCode: http.StatusNotFound,
			Type:       "mysql-error",
		}
	}

	return &result, nil
}

func (s UserSwipeRepository) Create(ctx *gin.Context, obj *models.UserSwipe) (*models.UserSwipe, *types.Error) {
	data := models.UserSwipe{}
	result, err := s.repository.Insert(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserSwipeStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	lastID, _ := (*result).LastInsertId()
	err = s.repository.FindByID(ctx, &data, lastID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserSwipeStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}
	return &data, nil
}

func (s UserSwipeRepository) Update(ctx *gin.Context, obj *models.UserSwipe) (*models.UserSwipe, *types.Error) {
	data := models.UserSwipe{}
	err := s.repository.Update(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserSwipeStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, obj.ID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserSwipeStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}
	return &data, nil
}
