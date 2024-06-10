package models

import (
	"anti-jomblo-go/library/types"
	"time"
)

type UserBulk struct {
	ID                 string    `json:"ID" db:"id"`
	Name               string    `json:"Name" db:"name"`
	Email              string    `json:"Email" db:"email"`
	CountryCallingCode string    `json:"CountryCallingCode" db:"country_calling_code"`
	PhoneNumber        string    `json:"PhoneNumber" db:"phone_number"`
	Password           string    `json:"Password" db:"password"`
	Gender             int       `json:"Gender" db:"gender"`
	BirthDate          time.Time `json:"BirthDate" db:"birth_date"`
	Height             int       `json:"Height" db:"height"`
	AboutMe            string    `json:"AboutMe" db:"about_me"`

	StatusID   string `json:"StatusID" db:"status_id"`
	StatusName string `json:"StatusName" db:"status_name"`
}

type User struct {
	ID                 string    `json:"ID" db:"id"`
	Name               string    `json:"Name" db:"name" validate:"required"`
	Email              string    `json:"Email" db:"email" validate:"required"`
	CountryCallingCode string    `json:"CountryCallingCode" db:"country_calling_code" validate:"required"`
	PhoneNumber        string    `json:"PhoneNumber" db:"phone_number" validate:"required"`
	Password           string    `json:"Password" db:"password" validate:"required"`
	Gender             int       `json:"Gender" db:"gender"`
	BirthDate          time.Time `json:"BirthDate" db:"birth_date" validate:"required"`
	Height             int       `json:"Height" db:"height"`
	AboutMe            string    `json:"AboutMe" db:"about_me"`

	StatusID string `json:"StatusID" db:"status_id"`
	Status   Status `json:"Status"`

	Pictures []*UserPicture `json:"Pictures"`
}

type UserJWTContent struct {
	ID    string `json:"ID" db:"id"`
	Name  string `json:"Name" db:"name" validate:"required"`
	Token string `json:"Token"`
	Email string `json:"Email" db:"email" validate:"required"`

	StatusID string `json:"StatusID" db:"status_id"`
	Status   Status `json:"Status"`
}

type FindAllUserParams struct {
	FindAllParams      types.FindAllParams
	Name               string
	Email              string
	CountryCallingCode string
	PhoneNumber        string
	Password           string
}
