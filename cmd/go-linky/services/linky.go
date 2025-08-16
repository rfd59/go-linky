package services

import (
	"fmt"
	"rfd59/go-linky/cmd/go-linky/core/linky"
	"rfd59/go-linky/cmd/go-linky/infra"
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/utils"

	"go.bug.st/serial"
)

type ILinkyService interface {
	ReadTic(frame []byte, mode linky.ILinkyMode) (*models.TiC, error)
	OpenPort(settings *models.Serial, serial infra.ISerialInfra) (serial.Port, error)
}

type LinkyService struct{}

func (s *LinkyService) ReadTic(frame []byte, mode linky.ILinkyMode) (*models.TiC, error) {
	// Check if the frame is valid
	if _, err := isValid(frame); err != nil {
		return nil, err
	}

	ds := mode.LoadDatasets(string(frame[1 : len(frame)-1])) // Exclude STX and ETX characters
	if len(ds) == 0 {
		return nil, utils.ErrNoDatasets
	}

	return mode.LoadTiC(ds), nil
}

func isValid(frame []byte) (bool, error) {
	if len(frame) == 0 {
		return false, utils.ErrEmptyFrame
	}
	if frame[0] != 0x02 {
		return false, utils.ErrSTXFrame
	}
	if frame[len(frame)-1] != 0x03 {
		return false, utils.ErrETXFrame
	}
	return true, nil
}

func (s *LinkyService) OpenPort(settings *models.Serial, serial infra.ISerialInfra) (serial.Port, error) {
	// Open the serial port with the specified mode
	port, err := serial.Open(settings.Port, settings.Mode)
	if err != nil {
		return nil, fmt.Errorf("failed to open the serial port %q: %w", settings.Port, err)
	}

	return port, nil
}
