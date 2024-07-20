package mailpit

type Mailpit struct {
	store store
}

func NewMailPit(pool *pgxpool.Pool) Mailpit {
	return Mailpit{pgstore.New(pool)}
}

func (mp Mailpit) SendConfirmTripEmailToTripOwner(trupID uuid.UUID) error {
	ctx := context.Background()
	trip, err := mp.store.GetTrip(ctx, trupID)
	if err != nil {
		return fmt.Errorf("mailpit: failed to get trip for SendConfirmTripEmailToTripOwner: %w", err)
	}