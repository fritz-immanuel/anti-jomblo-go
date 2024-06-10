package models

import (
	"anti-jomblo-go/library/types"
)

type PremiumPackageBulk struct {
	ID          string  `json:"ID" db:"id"`
	Name        string  `json:"Name" db:"name"`
	Description string  `json:"Description" db:"description"`
	Price       float64 `json:"Price" db:"price"`
	Duration    int     `json:"Duration" db:"duration"`

	StatusID   string `json:"StatusID" db:"status_id"`
	StatusName string `json:"StatusName" db:"status_name"`
}

type PremiumPackage struct {
	ID          string  `json:"ID" db:"id"`
	Name        string  `json:"Name" db:"name"`
	Description string  `json:"Description" db:"description"`
	Price       float64 `json:"Price" db:"price"`
	Duration    int     `json:"Duration" db:"duration"`

	StatusID string `json:"StatusID" db:"status_id"`
	Status   Status `json:"Status"`
}

type FindAllPremiumPackageParams struct {
	FindAllParams types.FindAllParams
}
