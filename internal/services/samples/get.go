package samples

import (
	"context"
	"fmt"
)

func (s *Service) Get(ctx context.Context, filterMap map[string]string) ([]*Sample, error) {
	repoSamples, err := s.samplesRepo.Get(ctx, filterMap)
	if err != nil {
		return nil, fmt.Errorf("samplesRepo.Get.error: %w", err)
	}

	samples := make([]*Sample, len(repoSamples))
	for idx, repoSample := range repoSamples {
		samples[idx] = &Sample{
			Text: repoSample.Text,
			Number: &BigNumber{
				Value: repoSample.Number,
			},
			Found: repoSample.Found,
			Type:  repoSample.Type,
		}
	}

	return samples, nil
}
