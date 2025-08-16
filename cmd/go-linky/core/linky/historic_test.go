package linky_test

import (
	"log/slog"
	"rfd59/go-linky/cmd/go-linky/core/linky"
	"rfd59/go-linky/cmd/go-linky/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thejerf/slogassert"
)

func TestHistoric_LoadDatasets_Valid(t *testing.T) {
	assert := assert.New(t)

	my := &linky.Historic{}

	// Test the LoadDatasets function
	data := my.LoadDatasets("\nADCO 811775653238 O\r\nOPTARIF BBR( S\r")

	// Assert the expected behavior
	assert.Len(data, 2, "Expected 2 datasets to be loaded")
	assert.Equal("ADCO", data[0].Label, "First dataset label should be 'ADCO'")
	assert.Equal("811775653238", data[0].Data, "First dataset data should be '811775653238'")
	assert.Equal(byte('O'), data[0].Checksum, "First dataset checksum should be 'O'")
	assert.Equal("OPTARIF", data[1].Label, "Second dataset label should be 'OPTARIF'")
	assert.Equal("BBR(", data[1].Data, "Second dataset data should be 'BBR('")
	assert.Equal(byte('S'), data[1].Checksum, "Second dataset checksum should be 'S'")
}

func TestHistoric_LoadDatasets_Invalid(t *testing.T) {
	assert := assert.New(t)

	my := &linky.Historic{}

	// Test the LoadDatasets function
	data := my.LoadDatasets("\nADCO811775653238 O\r\nADCO 81177 5653238 O\r\nOPTARIF BBR( fS\r")

	// Assert the expected behavior
	assert.Empty(data, "Expected no datasets to be loaded due to invalid format")
}

func TestHistoric_LoadTIC_Valid(t *testing.T) {
	assert := assert.New(t)

	my := &linky.Historic{}
	ds := []models.LinkyDataset{{Label: "ADCO", Data: "811775653238", Checksum: 'O'}, {Label: "PAPP", Data: "00280", Checksum: '+'}}

	// Test the LoadDatasets function
	tic := my.LoadTiC(ds)

	// Assert the expected behavior
	assert.Equal("811775653238", tic.ADCO)
	assert.Empty(tic.OPTARIF)
	assert.Equal(uint8(0), tic.ISOUSC)
	assert.Equal(uint32(0), tic.BASE)
	assert.Equal(uint32(0), tic.HCHC)
	assert.Equal(uint32(0), tic.HCHP)
	assert.Equal(uint32(0), tic.EJPHN)
	assert.Equal(uint32(0), tic.EJPHPM)
	assert.Equal(uint32(0), tic.BBRHCJB)
	assert.Equal(uint32(0), tic.BBRHPJB)
	assert.Equal(uint32(0), tic.BBRHCJW)
	assert.Equal(uint32(0), tic.BBRHPJW)
	assert.Equal(uint32(0), tic.BBRHCJR)
	assert.Equal(uint32(0), tic.BBRHPJR)
	assert.Equal(uint8(0), tic.PEJP)
	assert.Empty(tic.PTEC)
	assert.Empty(tic.DEMAIN)
	assert.Equal(uint8(0), tic.IINST)
	assert.Equal(uint8(0), tic.ADPS)
	assert.Equal(uint8(0), tic.IMAX)
	assert.Equal(uint16(280), tic.PAPP)
	assert.Empty(tic.HHPHC)
	assert.Empty(tic.MOTDETAT)
}

func TestHistoric_LoadTIC_InvalidChecksum(t *testing.T) {
	assert := assert.New(t)

	my := &linky.Historic{}
	ds := []models.LinkyDataset{{Label: "ADCO", Data: "811785653238", Checksum: 'O'}, {Label: "BBRHPJB", Data: "002110855", Checksum: '@'}, {Label: "PAPX", Data: "00280", Checksum: '+'}}

	// Testing handler for slog.
	log := slogassert.New(t, slog.LevelError, nil)
	logger := slog.New(log)
	slog.SetDefault(logger)

	// Test the LoadDatasets function
	tic := my.LoadTiC(ds)

	// Assert the expected behavior
	assert.Empty(tic.ADCO)
	assert.Empty(tic.OPTARIF)
	assert.Equal(uint8(0), tic.ISOUSC)
	assert.Equal(uint32(0), tic.BASE)
	assert.Equal(uint32(0), tic.HCHC)
	assert.Equal(uint32(0), tic.HCHP)
	assert.Equal(uint32(0), tic.EJPHN)
	assert.Equal(uint32(0), tic.EJPHPM)
	assert.Equal(uint32(0), tic.BBRHCJB)
	assert.Equal(uint32(2110855), tic.BBRHPJB)
	assert.Equal(uint32(0), tic.BBRHCJW)
	assert.Equal(uint32(0), tic.BBRHPJW)
	assert.Equal(uint32(0), tic.BBRHCJR)
	assert.Equal(uint32(0), tic.BBRHPJR)
	assert.Equal(uint8(0), tic.PEJP)
	assert.Empty(tic.PTEC)
	assert.Empty(tic.DEMAIN)
	assert.Equal(uint8(0), tic.IINST)
	assert.Equal(uint8(0), tic.ADPS)
	assert.Equal(uint8(0), tic.IMAX)
	assert.Equal(uint16(0), tic.PAPP)
	assert.Empty(tic.HHPHC)
	assert.Empty(tic.MOTDETAT)
	log.AssertSomePrecise(slogassert.LogMessageMatch{Message: "Invalid checksum found for a dataset!", Level: slog.LevelError})
}
