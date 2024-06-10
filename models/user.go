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
	GenderID           int       `json:"GenderID" db:"gender_id"`
	BirthDate          time.Time `json:"BirthDate" db:"birth_date"`
	Height             int       `json:"Height" db:"height"`
	AboutMe            string    `json:"AboutMe" db:"about_me"`

	StatusID   string `json:"StatusID" db:"status_id"`
	StatusName string `json:"StatusName" db:"status_name"`

	Age        int    `json:"Age" db:"age"`
	GenderName string `json:"GenderName" db:"gender_name"`
}

type User struct {
	ID                 string    `json:"ID" db:"id"`
	Name               string    `json:"Name" db:"name" validate:"required"`
	Email              string    `json:"Email" db:"email" validate:"required"`
	CountryCallingCode string    `json:"CountryCallingCode" db:"country_calling_code" validate:"required"`
	PhoneNumber        string    `json:"PhoneNumber" db:"phone_number" validate:"required"`
	Password           string    `json:"Password" db:"password" validate:"required"`
	GenderID           int       `json:"GenderID" db:"gender_id" validate:"required"`
	BirthDate          time.Time `json:"BirthDate" db:"birth_date" validate:"required"`
	Height             int       `json:"Height" db:"height"`
	AboutMe            string    `json:"AboutMe" db:"about_me"`

	StatusID string `json:"StatusID" db:"status_id"`
	Status   Status `json:"Status"`

	Gender *INTIDNameTemplate `json:"Gender"`

	Pictures []*UserPicture `json:"Pictures"`
}

type UserUpdate struct {
	ID                 string    `json:"ID" db:"id"`
	Name               string    `json:"Name" db:"name" validate:"required"`
	Email              string    `json:"Email" db:"email" validate:"required"`
	CountryCallingCode string    `json:"CountryCallingCode" db:"country_calling_code" validate:"required"`
	PhoneNumber        string    `json:"PhoneNumber" db:"phone_number" validate:"required"`
	GenderID           int       `json:"GenderID" db:"gender_id" validate:"required"`
	BirthDate          time.Time `json:"BirthDate" db:"birth_date" validate:"required"`
	Height             int       `json:"Height" db:"height"`
	AboutMe            string    `json:"AboutMe" db:"about_me"`

	StatusID string `json:"StatusID" db:"status_id"`

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

type UserForDatingList struct {
	ID       string `json:"ID" db:"id"`
	Name     string `json:"Name" db:"name"`
	GenderID int    `json:"GenderID" db:"gender_id"`
	Age      int    `json:"Age" db:"age"`
	Height   int    `json:"Height" db:"height"`
	AboutMe  string `json:"AboutMe" db:"about_me"`

	Gender INTIDNameTemplate `json:"Gender"`
}

type UserUpdatePassword struct {
	ID                 string `json:"ID"`
	OldPassword        string `json:"OldPassword" validate:"required"`
	NewPassword        string `json:"NewPassword" validate:"required"`
	NewPasswordConfirm string `json:"NewPasswordConfirm" validate:"required"`
}

type FindAllUserParams struct {
	FindAllParams      types.FindAllParams
	UserID             string
	Name               string
	Email              string
	CountryCallingCode string
	PhoneNumber        string
	Password           string
}
