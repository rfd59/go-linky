package linky

import (
	"fmt"
	"log/slog"
	"rfd59/go-linky/cmd/go-linky/models"
	"strings"
)

type Historic struct {
}

func (h *Historic) LoadDatasets(frame string) (datasets []models.LinkyDataset) {
	slog.Debug("Loading the datasets from 'Historique' frame...")

	for _, dataset := range strings.Split(frame, "\r") {
		if dataset == "" {
			continue // Skip empty dataset
		}
		// Load the dataset
		ld, err := h.loadDataset(dataset)
		if err != nil {
			slog.Error("Dataset '" + dataset + "' can't be loaded! [" + err.Error() + "]")
		} else {
			datasets = append(datasets, ld)
		}
	}

	return datasets
}

func (h *Historic) LoadTiC(ds []models.LinkyDataset) *models.TiC {
	tic := &models.TiC{}

	for _, dataset := range ds {
		// Calculate the checksum for the dataset
		input := h.getChecksumControlString(dataset)
		checksum := checksum(input)

		// Compare the calculated checksum with the dataset's checksum
		if dataset.Checksum != checksum {
			slog.Error("Invalid checksum found for a dataset!", "input", string(input), "expected", dataset.Checksum, "checksum", checksum)
			slog.Warn("Dataset '" + dataset.Label + "' will be ignored!")
		} else {
			mapToLinky(tic, dataset)
		}
	}

	return tic
}

func (h *Historic) loadDataset(dataset string) (ld models.LinkyDataset, err error) {
	items := strings.Split(dataset, " ")

	if len(items) == 3 {
		ld.Label = items[0][1:] // Remove the '\n' character
		ld.Data = items[1]
		// Checksum should be a single byte
		if len(items[2]) == 1 {
			ld.Checksum = items[2][0]
		} else {
			err = fmt.Errorf("Invalid checksum length: %s", items[2])
		}
	} else {
		err = fmt.Errorf("Invalid dataset format: %s", dataset)
	}

	return ld, err
}

func (h *Historic) getChecksumControlString(dataset models.LinkyDataset) []byte {
	// Return the input for checksum calculation
	return []byte(dataset.Label + " " + dataset.Data)
}
