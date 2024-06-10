package models

import (
	"anti-jomblo-go/library/types"
	"time"
)

type UserSwipeBulk struct {
	ID            string `json:"ID" db:"id"`
	UserID        string `json:"UserID" db:"user_id"`
	DisplayUserID string `json:"DisplayUserID" db:"display_user_id"`
	ActionID      int    `json:"ActionID" db:"action_id"`

	DisplayUserName string `json:"DisplayUserName" db:"display_user_name"`
}

type UserSwipe struct {
	ID            string `json:"ID" db:"id"`
	UserID        string `json:"UserID" db:"user_id"`
	DisplayUserID string `json:"DisplayUserID" db:"display_user_id"`
	ActionID      int    `json:"ActionID" db:"action_id"`

	DisplayUser *IDNameTemplate `json:"DisplayUser"`
}

type FindAllUserSwipeParams struct {
	FindAllParams types.FindAllParams
	StartDate     *time.Time
	EndDate       *time.Time
	UserID        string
	DisplayUserID string
	ActionID      int
}
