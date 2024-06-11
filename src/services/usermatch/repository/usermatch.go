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

type UserMatchRepository struct {
	repository data.GenericStorage
}

func NewUserMatchRepository(repository data.GenericStorage) UserMatchRepository {
	return UserMatchRepository{repository: repository}
}

func (s UserMatchRepository) FindAll(ctx *gin.Context, params models.FindAllUserMatchParams) ([]*models.UserMatch, *types.Error) {
	data := []*models.UserMatch{}
	bulks := []*models.UserMatchBulk{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.FindAllParams.StatusID != "" {
		where += fmt.Sprintf(` AND user_matches.%s`, params.FindAllParams.StatusID)
	}

	if params.StartDate != nil && !params.StartDate.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_matches.created_at) >= DATE("%s")`, params.StartDate.Format(library.StrToDateFormat))
	}

	if params.EndDate != nil && !params.EndDate.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_matches.created_at) <= DATE("%s")`, params.EndDate.Format(library.StrToDateFormat))
	}

	if params.Date != nil && !params.Date.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_matches.created_at) = DATE("%s")`, params.Date.Format(library.StrToDateFormat))
	}

	if params.UserID != "" {
		where += ` AND user_matches.user_id = :user_id`
	}

	if params.DisplayUserID != "" {
		where += ` AND user_matches.display_user_id = :display_user_id`
	}

	if params.FindAllParams.SortBy != "" {
		where += fmt.Sprintf(` ORDER BY %s`, params.FindAllParams.SortBy)
	}

	if params.FindAllParams.Page > 0 && params.FindAllParams.Size > 0 {
		where += ` LIMIT :limit OFFSET :offset`
	}

	query := fmt.Sprintf(`
  SELECT
    user_matches.id, user_matches.user_id, user_matches.display_user_id, user_matches.status_id,
    status.name status_name, users.name display_user_name
  FROM user_matches
  JOIN status ON status.id = user_matches.status_id
  JOIN users ON users.id = user_matches.display_user_id
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
			Path:       ".UserMatchStorage->FindAll()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	for _, v := range bulks {
		obj := &models.UserMatch{
			ID:            v.ID,
			UserID:        v.UserID,
			DisplayUserID: v.DisplayUserID,
			DisplayUser: &models.IDNameTemplate{
				ID:   v.DisplayUserID,
				Name: v.DisplayUserName,
			},
		}
		data = append(data, obj)
	}

	return data, nil
}

func (s UserMatchRepository) Count(ctx *gin.Context, params models.FindAllUserMatchParams) (int, *types.Error) {
	bulks := []*models.UserMatchBulk{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.FindAllParams.StatusID != "" {
		where += fmt.Sprintf(` AND user_matches.%s`, params.FindAllParams.StatusID)
	}

	if params.StartDate != nil && !params.StartDate.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_matches.created_at) >= DATE("%s")`, params.StartDate.Format(library.StrToDateFormat))
	}

	if params.EndDate != nil && !params.EndDate.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_matches.created_at) <= DATE("%s")`, params.EndDate.Format(library.StrToDateFormat))
	}

	if params.Date != nil && !params.Date.IsZero() {
		where += fmt.Sprintf(` AND DATE(user_matches.created_at) = DATE("%s")`, params.Date.Format(library.StrToDateFormat))
	}

	if params.UserID != "" {
		where += ` AND user_matches.user_id = :user_id`
	}

	if params.DisplayUserID != "" {
		where += ` AND user_matches.display_user_id = :display_user_id`
	}

	query := fmt.Sprintf(`
  SELECT
    user_matches.id, user_matches.user_id, user_matches.display_user_id, user_matches.status_id,
    status.name status_name, users.name display_user_name
  FROM user_matches
  JOIN status ON status.id = user_matches.status_id
  JOIN users ON users.id = user_matches.display_user_id
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
			Path:       ".UserMatchStorage->Count()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	return len(bulks), nil
}

func (s UserMatchRepository) Find(ctx *gin.Context, id string) (*models.UserMatch, *types.Error) {
	result := models.UserMatch{}
	bulks := []*models.UserMatchBulk{}
	var err error

	query := `
  SELECT
    user_matches.id, user_matches.user_id, user_matches.display_user_id,
    users.name display_user_name
  FROM user_matches
  JOIN users ON users.id = user_matches.display_user_id
  WHERE users.id = :id`

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserMatchStorage->Find()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	if len(bulks) > 0 {
		v := bulks[0]
		result = models.UserMatch{
			ID:            v.ID,
			UserID:        v.UserID,
			DisplayUserID: v.DisplayUserID,
			DisplayUser: &models.IDNameTemplate{
				ID:   v.DisplayUserID,
				Name: v.DisplayUserName,
			},
		}
	} else {
		return nil, &types.Error{
			Path:       ".UserMatchStorage->Find()",
			Message:    "Data Not Found",
			Error:      data.ErrNotFound,
			StatusCode: http.StatusNotFound,
			Type:       "mysql-error",
		}
	}

	return &result, nil
}

func (s UserMatchRepository) Create(ctx *gin.Context, obj *models.UserMatch) (*models.UserMatch, *types.Error) {
	data := models.UserMatch{}
	_, err := s.repository.Insert(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserMatchStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, obj.ID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserMatchStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}
	return &data, nil
}

func (s UserMatchRepository) Update(ctx *gin.Context, obj *models.UserMatch) (*models.UserMatch, *types.Error) {
	data := models.UserMatch{}
	err := s.repository.Update(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserMatchStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, obj.ID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserMatchStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}
	return &data, nil
}

func (s UserMatchRepository) FindAllUserMatches(ctx *gin.Context, params models.FindAllUserMatchParams) ([]*models.UserMatchPersonalList, *types.Error) {
	result := []*models.UserMatchPersonalList{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.FindAllParams.StatusID != "" {
		where += fmt.Sprintf(` AND um.%s`, params.FindAllParams.StatusID)
	}

	if params.UserID != "" {
		where += ` AND (um.user_id = :user_id OR um.display_user_id = :user_id)`
	}

	if params.FindAllParams.SortBy != "" {
		where += fmt.Sprintf(` ORDER BY %s`, params.FindAllParams.SortBy)
	}

	if params.FindAllParams.Page > 0 && params.FindAllParams.Size > 0 {
		where += ` LIMIT :limit OFFSET :offset`
	}

	query := fmt.Sprintf(`
  SELECT
    um.id, IF (um.user_id = "%s", um.display_user_id, um.user_id) matched_user_id,
    IF (um.user_id = "%s",
      (SELECT name FROM users WHERE id = um.display_user_id),
      (SELECT name FROM users WHERE id = um.user_id)
    ) matched_user_name,
    um.status_id, um.created_at
  FROM user_matches um
  WHERE %s
  `, params.UserID, params.UserID, where)

	err = s.repository.SelectWithQuery(ctx, &result, query, map[string]interface{}{
		"limit":     params.FindAllParams.Size,
		"offset":    ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id": params.FindAllParams.StatusID,
		"user_id":   params.UserID,
	})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserMatchStorage->FindAllUserMatches()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	return result, nil
}

func (s UserMatchRepository) CountUserMatches(ctx *gin.Context, params models.FindAllUserMatchParams) (int, *types.Error) {
	result := []*models.UserMatchPersonalList{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.FindAllParams.StatusID != "" {
		where += fmt.Sprintf(` AND um.%s`, params.FindAllParams.StatusID)
	}

	if params.UserID != "" {
		where += ` AND (um.user_id = :user_id OR um.display_user_id = :user_id)`
	}

	if params.FindAllParams.SortBy != "" {
		where += fmt.Sprintf(` ORDER BY %s`, params.FindAllParams.SortBy)
	}

	if params.FindAllParams.Page > 0 && params.FindAllParams.Size > 0 {
		where += ` LIMIT :limit OFFSET :offset`
	}

	query := fmt.Sprintf(`
  SELECT
    um.id, IF (um.user_id = "%s", um.display_user_id, um.user_id) matched_user_id,
    IF (um.user_id = "%s",
      (SELECT name FROM users WHERE id = um.display_user_id),
      (SELECT name FROM users WHERE id = um.user_id)
    ) matched_user_name,
    um.status_id, um.created_at
  FROM user_matches um
  WHERE %s
  `, params.UserID, params.UserID, where)

	err = s.repository.SelectWithQuery(ctx, &result, query, map[string]interface{}{
		"limit":     params.FindAllParams.Size,
		"offset":    ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id": params.FindAllParams.StatusID,
		"user_id":   params.UserID,
	})
	if err != nil {
		return 0, &types.Error{
			Path:       ".UserMatchStorage->FindAllUserMatches()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	return len(result), nil
}
