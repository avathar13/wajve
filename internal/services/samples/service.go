package samples

import (
	"context"
	"fmt"

	"wajve/internal/configs"
	"wajve/internal/repositories/samples"
)

type Service struct {
	samplesRepo samples.RepositoryInterface
}

type ServiceInterface interface {
	Get(ctx context.Context, filterMap map[string]string) ([]*Sample, error)
	Populate(ctx context.Context, path string) error
}

func New(cfg *configs.Cfg) (*Service, error) {
	samplesRepo, err := samples.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("samplesRepo.New.error: %w", err)
	}

	return &Service{
		samplesRepo: samplesRepo,
	}, nil
}
