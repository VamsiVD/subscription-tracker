package subscriptions

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionsCreateRequest struct {
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Type      string    `json:"type"` // "weekly" | "monthly" | "yearly"
	Amount    int64     `json:"amount"`
	StartDate time.Time `json:"start_date"`
	Currency  string    `json:"currency"`
	Owner     string    `json:"owner_id"`
}
type SubscriptionsResponse struct {
	Uuid      uuid.UUID `db:"uuid" json:"uuid"`
	Name      string    `db:"name" json:"name"`
	Type      string    `db:"type" json:"type"` // "weekly" | "monthly" | "yearly"
	Amount    int64     `db:"amount" json:"amount"`
	StartDate time.Time `db:"start_date" json:"start_date"`
	Currency  string    `db:"currency" json:"currency"`
	Owner     string    `db:"owner_id" json:"owner_id"`
}

type SubscriptionsUpdateRequest struct {
	Name      string    `json:"name,omitempty"`
	Type      string    `json:"type,omitempty"` // "weekly" | "monthly" | "yearly"
	Amount    int64     `json:"amount,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	Currency  string    `json:"currency"`
}

type SubscriptionsDeleteRequest struct {
	Uuid string `json:"uuid"`
}
