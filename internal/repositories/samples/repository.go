package samples

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"wajve/internal/configs"
)

type Repository struct {
	mongoDB *mongo.Client
}

type RepositoryInterface interface {
	Get(ctx context.Context, filterMap map[string]string) ([]*Sample, error)
	PopulateDB(ctx context.Context, samples []*Sample) error
}

func New(cfg *configs.Cfg) (*Repository, error) {
	credentials := ""
	if cfg.MongoDB.Username != "" {
		credentials = fmt.Sprintf("%s:%s@", cfg.MongoDB.Username, cfg.MongoDB.Password)
	}

	uri := fmt.Sprintf(
		"%s://%s%s:%d/?connectTimeoutMS=%d&maxPoolSize=%d&w=%s",
		cfg.MongoDB.Protocol,
		credentials,
		cfg.MongoDB.Host,
		cfg.MongoDB.Port,
		cfg.MongoDB.ConnectTimeout,
		cfg.MongoDB.MaxPoolSize,
		cfg.MongoDB.WriteConcern,
	)

	clientOptions := options.Client().ApplyURI(uri)

	mongoDB, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect.error: %w, uri: %s", err, uri)
	}

	if err = mongoDB.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("mongo.Ping.error: %w", err)
	}

	return &Repository{
		mongoDB: mongoDB,
	}, nil
}
