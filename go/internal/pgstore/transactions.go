package pgstore

import (
	"context"
	"fmt"
	"journey/internal/api/spec"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (q *Queries) CreateTrip(
	ctx context.Context,
	pool *pgxpool.Pool,
	params spec.CreateTripRequest,
) (uuid.UUID, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to begin tx for CreateTrip: %w", err)
	}
}
