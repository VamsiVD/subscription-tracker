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

func (r *Respository) Getsubscriptions(ctx context.Context, Owner string) ([]SubscriptionsResponse, error) {
	var subs []SubscriptionsResponse

	err := r.db.SelectContext(ctx, &subs,
		`SELECT uuid, owner_id, name, type, amount, start_date, currency FROM subscriptions WHERE owner_id = $1`,
		Owner,
	)
	return subs, err
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

func (r *Respository) UpdateSubscription(ctx context.Context, req SubscriptionsUpdateRequest, Uuid string) (SubscriptionsResponse, error) {
	var sub SubscriptionsResponse
	err := r.db.GetContext(ctx, &sub,
		`UPDATE subscriptions 
			 SET name = $1, type = $2, amount = $3, start_date = $4, currency = $5
			 WHERE uuid = $6
			 RETURNING uuid, owner_id, name, type, amount, start_date, currency`,
		req.Name, req.Type, req.Amount, req.StartDate, req.Currency, Uuid,
	)
	return sub, err
}

func (r *Respository) DeleteSubscription(ctx context.Context, Uuid string) error {

	_, err := r.db.ExecContext(ctx,
		`DELETE FROM subscriptions 
			 WHERE uuid = $1`,
		Uuid,
	)

	return err
}

func generateUuid() (string, error) {
	id := uuid.New()
	return id.String(), nil
}
