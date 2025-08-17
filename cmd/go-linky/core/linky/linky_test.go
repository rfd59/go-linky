package linky_test

import (
	"rfd59/go-linky/cmd/go-linky/core/linky"
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinky_Historic_Tempo(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Test a true Tempo Frame
	service := services.LinkyService{}
	tic, err := service.ReadTic([]byte{2, 10, 65, 68, 67, 79, 32, 56, 49, 49, 55, 55, 53, 54, 53, 51, 50, 51, 56, 32, 79, 13, 10, 79, 80, 84, 65, 82, 73, 70, 32, 66, 66, 82, 40, 32, 83, 13, 10, 73, 83, 79, 85, 83, 67, 32, 51, 48, 32, 57, 13, 10, 66, 66, 82, 72, 67, 74, 66, 32, 48, 49, 53, 55, 52, 48, 49, 52, 57, 32, 60, 13, 10, 66, 66, 82, 72, 80, 74, 66, 32, 48, 48, 50, 49, 49, 55, 53, 51, 53, 32, 66, 13, 10, 66, 66, 82, 72, 67, 74, 87, 32, 48, 48, 48, 50, 52, 50, 57, 57, 51, 32, 79, 13, 10, 66, 66, 82, 72, 80, 74, 87, 32, 48, 48, 48, 52, 51, 54, 54, 49, 57, 32, 92, 13, 10, 66, 66, 82, 72, 67, 74, 82, 32, 48, 48, 48, 48, 56, 48, 55, 53, 48, 32, 65, 13, 10, 66, 66, 82, 72, 80, 74, 82, 32, 48, 48, 48, 48, 54, 55, 48, 54, 55, 32, 84, 13, 10, 80, 84, 69, 67, 32, 72, 80, 74, 66, 32, 80, 13, 10, 68, 69, 77, 65, 73, 78, 32, 45, 45, 45, 45, 32, 34, 13, 10, 73, 73, 78, 83, 84, 32, 48, 48, 48, 32, 87, 13, 10, 73, 77, 65, 88, 32, 48, 57, 48, 32, 72, 13, 10, 80, 65, 80, 80, 32, 48, 48, 50, 50, 48, 32, 37, 13, 10, 72, 72, 80, 72, 67, 32, 65, 32, 44, 13, 10, 77, 79, 84, 68, 69, 84, 65, 84, 32, 48, 48, 48, 48, 48, 48, 32, 66, 13, 3}, &linky.Historic{})

	// Assert the expected behavior
	require.NoError(err, "Expected no error to analyze frame")
	assert.Equal("811775653238", tic.ADCO)
	assert.Equal("BBR(", tic.OPTARIF)
	assert.Equal(uint8(30), tic.ISOUSC)
	assert.Equal(uint32(0), tic.BASE)
	assert.Equal(uint32(0), tic.HCHC)
	assert.Equal(uint32(0), tic.HCHP)
	assert.Equal(uint32(0), tic.EJPHN)
	assert.Equal(uint32(0), tic.EJPHPM)
	assert.Equal(uint32(15740149), tic.BBRHCJB)
	assert.Equal(uint32(2117535), tic.BBRHPJB)
	assert.Equal(uint32(242993), tic.BBRHCJW)
	assert.Equal(uint32(436619), tic.BBRHPJW)
	assert.Equal(uint32(80750), tic.BBRHCJR)
	assert.Equal(uint32(67067), tic.BBRHPJR)
	assert.Equal(uint8(0), tic.PEJP)
	assert.Equal("HPJB", tic.PTEC)
	assert.Equal("----", tic.DEMAIN)
	assert.Equal(uint8(0), tic.IINST)
	assert.Equal(uint8(0), tic.ADPS)
	assert.Equal(uint8(90), tic.IMAX)
	assert.Equal(uint16(220), tic.PAPP)
	assert.Equal("A", tic.HHPHC)
	assert.Equal("000000", tic.MOTDETAT)
}

func TestLinky_MapToLinky(t *testing.T) {
	assert := assert.New(t)

	my := &linky.Historic{}
	ds := []models.LinkyDataset{
		{Label: "BASE", Data: "12345", Checksum: 'Z'},
		{Label: "HCHC", Data: "12345", Checksum: 'U'},
		{Label: "HCHP", Data: "12345", Checksum: '"'},
		{Label: "EJPHN", Data: "12345", Checksum: '4'},
		{Label: "EJPHPM", Data: "12345", Checksum: 'C'},
		{Label: "PEJP", Data: "123", Checksum: 'E'},
		{Label: "ADPS", Data: "45", Checksum: 'Q'},
		{Label: "XXXX", Data: "Unrecognized", Checksum: 'M'},
	}

	// Test the LoadDatasets function
	tic := my.LoadTiC(ds)

	// Assert the expected behavior
	assert.Equal(uint32(12345), tic.BASE)
	assert.Equal(uint32(12345), tic.HCHC)
	assert.Equal(uint32(12345), tic.HCHP)
	assert.Equal(uint32(12345), tic.EJPHN)
	assert.Equal(uint32(12345), tic.EJPHPM)
	assert.Equal(uint8(123), tic.PEJP)
	assert.Equal(uint8(45), tic.ADPS)
}
