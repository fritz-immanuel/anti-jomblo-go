package repository

import (
	"fmt"
	"net/http"

	"anti-jomblo-go/library/data"
	"anti-jomblo-go/library/types"
	"anti-jomblo-go/models"

	"github.com/gin-gonic/gin"
)

type PremiumPackageRepository struct {
	repository       data.GenericStorage
	statusRepository data.GenericStorage
}

func NewPremiumPackageRepository(repository data.GenericStorage, statusRepository data.GenericStorage) PremiumPackageRepository {
	return PremiumPackageRepository{repository: repository, statusRepository: statusRepository}
}

func (s PremiumPackageRepository) FindAll(ctx *gin.Context, params models.FindAllPremiumPackageParams) ([]*models.PremiumPackage, *types.Error) {
	data := []*models.PremiumPackage{}
	bulks := []*models.PremiumPackageBulk{}

	var err error

	where := `TRUE`

	if params.FindAllParams.DataFinder != "" {
		where += fmt.Sprintf(` AND %s`, params.FindAllParams.DataFinder)
	}

	if params.FindAllParams.StatusID != "" {
		where += fmt.Sprintf(` AND premium_packages.%s`, params.FindAllParams.StatusID)
	}

	if params.FindAllParams.SortBy != "" {
		where += fmt.Sprintf(` ORDER BY %s`, params.FindAllParams.SortBy)
	}

	if params.FindAllParams.Page > 0 && params.FindAllParams.Size > 0 {
		where += ` LIMIT :limit OFFSET :offset`
	}

	query := fmt.Sprintf(`
  SELECT
    premium_packages.id, premium_packages.name, premium_packages.description, premium_packages.price, premium_packages.duration,
    premium_packages.status_id, status.name status_name
  FROM premium_packages
  JOIN status ON status.id = premium_packages.status_id
  WHERE %s
  `, where)

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{
		"limit":     params.FindAllParams.Size,
		"offset":    ((params.FindAllParams.Page - 1) * params.FindAllParams.Size),
		"status_id": params.FindAllParams.StatusID,
	})
	if err != nil {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->FindAll()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	for _, v := range bulks {
		obj := &models.PremiumPackage{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Duration:    v.Duration,
			StatusID:    v.StatusID,
			Status: models.Status{
				ID:   v.StatusID,
				Name: v.StatusName,
			},
		}
		data = append(data, obj)
	}

	return data, nil
}

func (s PremiumPackageRepository) Find(ctx *gin.Context, id string) (*models.PremiumPackage, *types.Error) {
	result := models.PremiumPackage{}
	bulks := []*models.PremiumPackageBulk{}
	var err error

	query := `
  SELECT
    premium_packages.id, premium_packages.name, premium_packages.description, premium_packages.price, premium_packages.duration,
    premium_packages.status_id, status.name status_name
  FROM premium_packages
  JOIN status ON status.id = premium_packages.status_id
  WHERE premium_packages.id = :id`

	err = s.repository.SelectWithQuery(ctx, &bulks, query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->Find()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	if len(bulks) > 0 {
		v := bulks[0]
		result = models.PremiumPackage{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Duration:    v.Duration,
			StatusID:    v.StatusID,
			Status: models.Status{
				ID:   v.StatusID,
				Name: v.StatusName,
			},
		}
	} else {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->Find()",
			Message:    "Data Not Found",
			Error:      data.ErrNotFound,
			StatusCode: http.StatusNotFound,
			Type:       "mysql-error",
		}
	}

	return &result, nil
}

func (s PremiumPackageRepository) Create(ctx *gin.Context, obj *models.PremiumPackage) (*models.PremiumPackage, *types.Error) {
	data := models.PremiumPackage{}
	_, err := s.repository.Insert(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, obj.ID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->Create()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}
	return &data, nil
}

func (s PremiumPackageRepository) Update(ctx *gin.Context, obj *models.PremiumPackage) (*models.PremiumPackage, *types.Error) {
	data := models.PremiumPackage{}
	err := s.repository.Update(ctx, obj)
	if err != nil {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, obj.ID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->Update()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}
	return &data, nil
}

func (s PremiumPackageRepository) FindStatus(ctx *gin.Context) ([]*models.Status, *types.Error) {
	status := []*models.Status{}

	err := s.statusRepository.Where(ctx, &status, "1=1", map[string]interface{}{})
	if err != nil {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->FindStatus()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	return status, nil
}

func (s PremiumPackageRepository) UpdateStatus(ctx *gin.Context, id string, statusID string) (*models.PremiumPackage, *types.Error) {
	data := models.PremiumPackage{}
	err := s.repository.UpdateStatus(ctx, id, statusID)
	if err != nil {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->UpdateStatus()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	err = s.repository.FindByID(ctx, &data, id)
	if err != nil {
		return nil, &types.Error{
			Path:       ".PremiumPackageStorage->UpdateStatus()",
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Type:       "mysql-error",
		}
	}

	return &data, nil
}
