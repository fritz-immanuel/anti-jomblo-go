package models

import (
	"anti-jomblo-go/library/types"
	"time"
)

type UserMatchBulk struct {
	ID            string `json:"ID" db:"id"`
	UserID        string `json:"UserID" db:"user_id"`
	DisplayUserID string `json:"DisplayUserID" db:"display_user_id"`

	StatusID   string `json:"StatusID" db:"status_id"`
	StatusName string `json:"StatusName" db:"status_name"`

	DisplayUserName string `json:"DisplayUserName" db:"display_user_name"`
}

type UserMatch struct {
	ID            string `json:"ID" db:"id"`
	UserID        string `json:"UserID" db:"user_id"`
	DisplayUserID string `json:"DisplayUserID" db:"display_user_id"`
	StatusID      string `json:"StatusID" db:"status_id"`

	Status      *IDNameTemplate `json:"Status"`
	DisplayUser *IDNameTemplate `json:"DisplayUser"`
}

type UserMatchPersonalList struct {
	ID              string    `json:"ID" db:"id"`
	MatchedUserID   string    `json:"MatchedUserID" db:"matched_user_id"`
	MatchedUserName string    `json:"MatchedUserName" db:"matched_user_name"`
	StatusID        string    `json:"StatusID" db:"status_id"`
	CreatedAt       time.Time `json:"CreatedAt" db:"created_at"`
}

type FindAllUserMatchParams struct {
	FindAllParams types.FindAllParams
	StartDate     *time.Time
	EndDate       *time.Time
	Date          *time.Time
	UserID        string
	DisplayUserID string
}
