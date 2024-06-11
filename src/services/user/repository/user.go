package repository

import (
	"fmt"
	"net/http"

	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	repository       data.GenericStorage
	statusRepository data.GenericStorage
}

func NewUserRepository(repository data.GenericStorage, statusRepository data.GenericStorage) UserRepository {
	return UserRepository{repository: repository, statusRepository: statusRepository}
}

func (s UserRepository) FindAll(ctx *gin.Context, params models.FindAllUserParams) ([]*models.User, *types.Error) {
	data := []*models.User{}
	bulks := []*models.UserBulk{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.FindAllParams.StatusID != "" {
		where += fmt.Sprintf(` AND users.%s`, params.FindAllParams.StatusID)
	}

	if params.Name != "" {
		where += fmt.Sprintf(` AND users.name LIKE "%s%%"`, params.Name)
	}

	if params.Email != "" {
		where += ` AND users.email = :email`
	}

	if params.Password != "" {
		where += ` AND users.password = :password`
	}

	if params.CountryCallingCode != "" {
		where += ` AND users.country_calling_code = :country_calling_code`
	}

	if params.PhoneNumber != "" {
		where += ` AND users.phone_number = :phone_number`
	}

	if params.FindAllParams.SortBy != "" {
		where += fmt.Sprintf(` ORDER BY %s`, params.FindAllParams.SortBy)
	}

	if params.FindAllParams.Page > 0 && params.FindAllParams.Size > 0 {
		where += ` LIMIT :limit OFFSET :offset`
	}

	query := fmt.Sprintf(`
  SELECT
    users.id, users.name, users.email, users.country_calling_code, users.phone_number,
    users.password, users.gender_id, users.birth_date, users.height, users.about_me,
    users.status_id, status.name status_name, genders.name gender_name
  FROM users
  JOIN status ON users.status_id = status.id
  JOIN genders ON genders.id = users.gender_id
  WHERE %s
  `, where)

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{
		"limit":                params.FindAllParams.Size,
		"offset":               ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id":            params.FindAllParams.StatusID,
		"email":                params.Email,
		"password":             params.Password,
		"country_calling_code": params.CountryCallingCode,
		"phone_number":         params.PhoneNumber,
	})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->FindAll()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	for _, v := range bulks {
		obj := &models.User{
			ID:                 v.ID,
			Name:               v.Name,
			Email:              v.Email,
			CountryCallingCode: v.CountryCallingCode,
			PhoneNumber:        v.PhoneNumber,
			Password:           v.Password,
			GenderID:           v.GenderID,
			Gender: &models.INTIDNameTemplate{
				ID:   v.GenderID,
				Name: v.GenderName,
			},
			BirthDate: v.BirthDate,
			Height:    v.Height,
			AboutMe:   v.AboutMe,
			StatusID:  v.StatusID,
			Status: models.Status{
				ID:   v.StatusID,
				Name: v.StatusName,
			},
		}
		data = append(data, obj)
	}

	return data, nil
}

func (s UserRepository) Find(ctx *gin.Context, id string) (*models.User, *types.Error) {
	result := models.User{}
	bulks := []*models.UserBulk{}
	var err error

	query := `
  SELECT
    users.id, users.name, users.email, users.country_calling_code, users.phone_number,
    users.password, users.gender_id, users.birth_date, users.height, users.about_me,
    users.status_id, status.name status_name, genders.name gender_name
  FROM users
  JOIN status ON users.status_id = status.id
  JOIN genders ON genders.id = users.gender_id
  WHERE users.id = :id`

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->Find()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	if len(bulks) > 0 {
		v := bulks[0]
		result = models.User{
			ID:                 v.ID,
			Name:               v.Name,
			Email:              v.Email,
			CountryCallingCode: v.CountryCallingCode,
			PhoneNumber:        v.PhoneNumber,
			Password:           v.Password,
			GenderID:           v.GenderID,
			Gender: &models.INTIDNameTemplate{
				ID:   v.GenderID,
				Name: v.GenderName,
			},
			BirthDate: v.BirthDate,
			Height:    v.Height,
			AboutMe:   v.AboutMe,
			StatusID:  v.StatusID,
			Status: models.Status{
				ID:   v.StatusID,
				Name: v.StatusName,
			},
		}
	} else {
		return nil, &types.Error{
			Path:       ".UserStorage->Find()",
			Message:    "Data Not Found",
			Error:      data.ErrNotFound,
			StatusCode: http.StatusNotFound,
			Type:       "mysql-error",
		}
	}

	return &result, nil
}

func (s UserRepository) Count(ctx *gin.Context, params models.FindAllUserParams) (int, *types.Error) {
	bulks := []*models.UserBulk{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.FindAllParams.StatusID != "" {
		where += fmt.Sprintf(` AND users.%s`, params.FindAllParams.StatusID)
	}

	if params.Name != "" {
		where += fmt.Sprintf(` AND users.name LIKE "%s%%"`, params.Name)
	}

	if params.Email != "" {
		where += ` AND users.email = :email`
	}

	if params.Password != "" {
		where += ` AND users.password = :password`
	}

	if params.CountryCallingCode != "" {
		where += ` AND users.country_calling_code = :country_calling_code`
	}

	if params.PhoneNumber != "" {
		where += ` AND users.phone_number = :phone_number`
	}

	query := fmt.Sprintf(`
  SELECT
    users.id, users.name, users.email, users.country_calling_code, users.phone_number,
    users.password, users.gender_id, users.birth_date, users.height, users.about_me,
    users.status_id, status.name status_name, genders.name gender_name
  FROM users
  JOIN status ON users.status_id = status.id
  JOIN genders ON genders.id = users.gender_id
  WHERE %s
  `, where)

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{
		"limit":                params.FindAllParams.Size,
		"offset":               ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id":            params.FindAllParams.StatusID,
		"email":                params.Email,
		"password":             params.Password,
		"country_calling_code": params.CountryCallingCode,
		"phone_number":         params.PhoneNumber,
	})
	if err != nil {
		return 0, &types.Error{
			Path:       ".UserStorage->Count()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	return len(bulks), nil
}

func (s UserRepository) Create(ctx *gin.Context, obj *models.User) (*models.User, *types.Error) {
	data := models.User{}
	_, err := s.repository.Insert(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, obj.ID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}
	return &data, nil
}

func (s UserRepository) Update(ctx *gin.Context, obj *models.User) (*models.User, *types.Error) {
	data := models.User{}
	err := s.repository.Update(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, obj.ID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}
	return &data, nil
}

func (s UserRepository) FindStatus(ctx *gin.Context) ([]*models.Status, *types.Error) {
	status := []*models.Status{}

	err := s.statusRepository.Where(ctx, &status, "1=1", map[string]interface{}{})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->FindStatus()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	return status, nil
}

func (s UserRepository) UpdateStatus(ctx *gin.Context, id string, statusID string) (*models.User, *types.Error) {
	data := models.User{}
	err := s.repository.UpdateStatus(ctx, id, statusID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->UpdateStatus()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, id)
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->UpdateStatus()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	return &data, nil
}

func (s UserRepository) FindAllForDating(ctx *gin.Context, params models.FindAllUserParams) ([]*models.UserForDatingList, *types.Error) {
	data := []*models.UserForDatingList{}
	bulks := []*models.UserBulk{}

	var err error

	where := `TRUE AND users.status_id = "1"`

	if params.UserID != "" {
		where += fmt.Sprintf(` AND users.id != "%s"`, params.UserID)
	}

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.FindAllParams.StatusID != "" {
		where += fmt.Sprintf(` AND users.%s`, params.FindAllParams.StatusID)
	}

	where += ` ORDER BY RAND()`

	if params.FindAllParams.Page > 0 && params.FindAllParams.Size > 0 {
		where += ` LIMIT :limit OFFSET :offset`
	}

	query := fmt.Sprintf(`
  SELECT
    users.id, users.name, users.gender_id,
    TIMESTAMPDIFF(YEAR, users.birth_date, UTC_TIMESTAMP + INTERVAL 7 HOUR) age,
    users.height, users.about_me,
    genders.name gender_name
  FROM users
  JOIN genders ON genders.id = users.gender_id
  LEFT JOIN user_swipes ON user_swipes.display_user_id = users.id
    AND DATE(user_swipes.created_at) = DATE(UTC_TIMESTAMP + INTERVAL 7 HOUR)
    AND user_swipes.user_id = "%s"
    AND user_swipes.id IS NULL
  WHERE %s
  `, params.UserID, where)

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{
		"limit":     params.FindAllParams.Size,
		"offset":    ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id": params.FindAllParams.StatusID,
	})
	if err != nil {
		return nil, &types.Error{
			Path:       ".UserStorage->FindAllForDating()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	for _, v := range bulks {
		obj := &models.UserForDatingList{
			ID:      v.ID,
			Name:    v.Name,
			Age:     v.Age,
			Height:  v.Height,
			AboutMe: v.AboutMe,
			Gender: models.INTIDNameTemplate{
				ID:   v.GenderID,
				Name: v.GenderName,
			},
		}

		data = append(data, obj)
	}

	return data, nil
}

func (s UserRepository) CountForDating(ctx *gin.Context, params models.FindAllUserParams) (int, *types.Error) {
	bulks := []*models.UserBulk{}

	var err error

	where := `TRUE AND users.status_id = "1"`

	if params.UserID != "" {
		where += fmt.Sprintf(` AND users.id != "%s"`, params.UserID)
	}

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.FindAllParams.StatusID != "" {
		where += fmt.Sprintf(` AND users.%s`, params.FindAllParams.StatusID)
	}

	where += ` ORDER BY RAND()`

	if params.FindAllParams.Page > 0 && params.FindAllParams.Size > 0 {
		where += ` LIMIT :limit OFFSET :offset`
	}

	query := fmt.Sprintf(`
  SELECT
    users.id, users.name, users.gender_id,
    TIMESTAMPDIFF(YEAR, users.birth_date, UTC_TIMESTAMP + INTERVAL 7 HOUR) age,
    users.height, users.about_me,
    genders.name gender_name
  FROM users
  JOIN genders ON genders.id = users.gender_id
  LEFT JOIN user_swipes ON user_swipes.display_user_id = users.id
    AND DATE(user_swipes.created_at) = DATE(UTC_TIMESTAMP + INTERVAL 7 HOUR)
    AND user_swipes.user_id = "%s"
    AND user_swipes.id IS NULL
  WHERE %s
  `, params.UserID, where)

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{
		"limit":     params.FindAllParams.Size,
		"offset":    ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id": params.FindAllParams.StatusID,
	})
	if err != nil {
		return 0, &types.Error{
			Path:       ".UserStorage->FindAll()",
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
