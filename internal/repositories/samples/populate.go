package samples

import (
	"context"
	"fmt"
)

func (r *Repository) PopulateDB(ctx context.Context, samples []*Sample) error {
	collection := r.mongoDB.Database("wajve").Collection("samples")

	dataToInsert := make([]interface{}, len(samples))
	for idx, sample := range samples {
		dataToInsert[idx] = sample
	}

	if _, err := collection.InsertMany(ctx, dataToInsert); err != nil {
		return fmt.Errorf("collection.InsertMany.error: %w", err)
	}

	return nil
}
