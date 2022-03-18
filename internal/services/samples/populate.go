package samples

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	samplesRepo "wajve/internal/repositories/samples"
)

func (s *Service) Populate(ctx context.Context, path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("os.Open.error: %w", err)
	}

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll.error: %w", err)
	}

	var samples []*Sample

	err = json.Unmarshal(jsonBytes, &samples)
	if err != nil {
		return fmt.Errorf("json.Unmarshal.error: %w", err)
	}

	repoSamples := make([]*samplesRepo.Sample, len(samples))

	for idx, sample := range samples {
		number := ""
		if sample.Number != nil {
			number = sample.Number.Value
		}

		repoSamples[idx] = &samplesRepo.Sample{
			Text:   sample.Text,
			Number: number,
			Found:  sample.Found,
			Type:   sample.Type,
		}
	}

	if err = s.samplesRepo.PopulateDB(ctx, repoSamples); err != nil {
		return fmt.Errorf("samplesRepo.PopulateDB.error: %w", err)
	}

	return nil
}
