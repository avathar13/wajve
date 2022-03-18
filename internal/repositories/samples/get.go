package samples

import (
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

var errNotSupportedFieldName = fmt.Errorf("this field name is not supported")

func (r *Repository) Get(ctx context.Context, filterMap map[string]string) ([]*Sample, error) {
	collection := r.mongoDB.Database("wajve").Collection("samples")

	bsonArr := make([]bson.M, 0)

	for fieldName, fieldValue := range filterMap {
		value, err := r.buildFilterValue(fieldName, fieldValue)
		if err != nil {
			return nil, fmt.Errorf("buildFilterValue: %w", err)
		}

		bsonArr = append(bsonArr, bson.M{fieldName: value})
	}

	filter := bson.M{
		"$and": bsonArr,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("mongoDB.collection.Find.error: %w", err)
	}

	var samples []*Sample
	if err = cursor.All(ctx, &samples); err != nil {
		return nil, fmt.Errorf("mongoDB.cursor.All.error: %w", err)
	}

	return samples, nil
}

func (r *Repository) buildFilterValue(fieldName string, fieldValue string) (interface{}, error) {
	switch fieldName {
	case "found":
		value, err := strconv.ParseBool(fieldValue)
		if err != nil {
			return nil, fmt.Errorf("strconv.ParseBool.error: %w, value: %s", err, fieldValue)
		}

		return value, nil
	case "text", "number", "type":
		return fieldValue, nil
	default:
		return nil, fmt.Errorf("%w: %s", errNotSupportedFieldName, fieldName)
	}
}
