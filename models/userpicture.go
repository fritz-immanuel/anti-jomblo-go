package models

import (
	"anti-jomblo-go/library/types"
)

type UserPictureBulk struct {
	ID     string `json:"ID" db:"id"`
	UserID string `json:"UserID" db:"user_id"`
	ImgURL string `json:"ImgURL" db:"img_url"`
	IsMain int    `json:"IsMain" db:"is_main"`
}

type UserPicture struct {
	ID     string `json:"ID" db:"id"`
	UserID string `json:"UserID" db:"user_id"`
	ImgURL string `json:"ImgURL" db:"img_url"`
	IsMain int    `json:"IsMain" db:"is_main"`
}

type FindAllUserPictureParams struct {
	FindAllParams types.FindAllParams
	UserID        string
	IsMain        int
}
