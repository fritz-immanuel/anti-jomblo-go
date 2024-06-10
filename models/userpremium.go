package models

import (
	"anti-jomblo-go/library/types"
	"time"
)

type UserPremiumBulk struct {
	ID        string    `json:"ID" db:"id"`
	UserID    string    `json:"UserID" db:"user_id"`
	BoughtAt  time.Time `json:"BoughtAt" db:"bought_at"`
	ExpiredAt time.Time `json:"ExpiredAt" db:"expired_at"`
}

type UserPremium struct {
	ID        string    `json:"ID" db:"id"`
	UserID    string    `json:"UserID" db:"user_id"`
	BoughtAt  time.Time `json:"BoughtAt" db:"bought_at"`
	ExpiredAt time.Time `json:"ExpiredAt" db:"expired_at"`
}

type FindAllUserPremiumParams struct {
	FindAllParams types.FindAllParams
	UserID        string
	NotExpired    int
}
