package subscriptions

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

var ErrAccountNotFound = errors.New("account not found")

type Respository struct {
	db *sqlx.DB
}

func NewRespository(db *sqlx.DB) *Respository {
	return &Respository{db: db}
}

func (r *Respository) CreateSubscription(ctx context.Context, req SubscriptionsCreateRequest) (SubscriptionsResponse, error) {
	var sub SubscriptionsResponse
	uuid, err := generateUuid()
	if err != nil {
		return SubscriptionsResponse{}, err
	}

	err = r.db.GetContext(ctx, &sub,
		`INSERT INTO subscriptions (uuid, owner_id, name, type, amount, start_date, currency)
				 VALUES ($1, $2, $3, $4, $5, $6, $7)
				 RETURNING uuid, owner_id, name, type, amount, start_date, currency`,
		uuid, req.Owner, req.Name, req.Type, req.Amount, req.StartDate, req.Currency,
	)
	return sub, err
}

func generateUuid() (string, error) {
	id := uuid.New()
	return id.String(), nil
}
