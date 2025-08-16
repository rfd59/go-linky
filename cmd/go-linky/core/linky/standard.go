package linky

import (
	"log/slog"
	"rfd59/go-linky/cmd/go-linky/models"
)

type Standard struct {
}

func (h *Standard) LoadDatasets(frame string) (datasets []models.LinkyDataset) {
	slog.Debug("Loading the datasets from 'Standard' frame...")
	panic("unimplemented")
}

func (h *Standard) LoadTiC(ds []models.LinkyDataset) *models.TiC {
	panic("unimplemented")
}
