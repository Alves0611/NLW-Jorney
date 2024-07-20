package pgstore

import (
	"context"
	"journey/internal/api/spec"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (q *Queries) CreateTrip(
	ctx context.Context,
	pool *pgxpool.Pool,
	params spec.CreateTripRequest,
) (uuid.UUID, error) {}
