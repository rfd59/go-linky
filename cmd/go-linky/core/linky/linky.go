package linky

import (
	"log/slog"
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/utils"
)

type ILinkyMode interface {
	LoadDatasets(frame string) []models.LinkyDataset
	LoadTiC(ds []models.LinkyDataset) *models.TiC
}

func checksum(control []byte) byte {
	var sum byte
	for _, b := range control {
		sum += b
	}
	return (sum & 0x3F) + 0x20
}

func mapToLinky(tic *models.TiC, dataset models.LinkyDataset) {
	// Map the dataset to the Linky model
	switch dataset.Label {
	case "ADCO":
		tic.ADCO = dataset.Data
	case "OPTARIF":
		tic.OPTARIF = dataset.Data
	case "ISOUSC":
		tic.ISOUSC = utils.ParseUint8(dataset.Data)
	case "BASE":
		tic.BASE = utils.ParseUint32(dataset.Data)
	case "HCHC":
		tic.HCHC = utils.ParseUint32(dataset.Data)
	case "HCHP":
		tic.HCHP = utils.ParseUint32(dataset.Data)
	case "EJPHN":
		tic.EJPHN = utils.ParseUint32(dataset.Data)
	case "EJPHPM":
		tic.EJPHPM = utils.ParseUint32(dataset.Data)
	case "BBRHCJB":
		tic.BBRHCJB = utils.ParseUint32(dataset.Data)
	case "BBRHPJB":
		tic.BBRHPJB = utils.ParseUint32(dataset.Data)
	case "BBRHCJW":
		tic.BBRHCJW = utils.ParseUint32(dataset.Data)
	case "BBRHPJW":
		tic.BBRHPJW = utils.ParseUint32(dataset.Data)
	case "BBRHCJR":
		tic.BBRHCJR = utils.ParseUint32(dataset.Data)
	case "BBRHPJR":
		tic.BBRHPJR = utils.ParseUint32(dataset.Data)
	case "PEJP":
		tic.PEJP = utils.ParseUint8(dataset.Data)
	case "PTEC":
		tic.PTEC = dataset.Data
	case "DEMAIN":
		tic.DEMAIN = dataset.Data
	case "IINST":
		tic.IINST = utils.ParseUint8(dataset.Data)
	case "ADPS":
		tic.ADPS = utils.ParseUint8(dataset.Data)
	case "IMAX":
		tic.IMAX = utils.ParseUint8(dataset.Data)
	case "PAPP":
		tic.PAPP = utils.ParseUint16(dataset.Data)
	case "HHPHC":
		tic.HHPHC = dataset.Data
	case "MOTDETAT":
		tic.MOTDETAT = dataset.Data
	default:
		slog.Warn("Unrecognized dataset label '" + dataset.Label + "'")
	}
}
