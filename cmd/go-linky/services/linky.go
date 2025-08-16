package services

import (
	"errors"
	"fmt"
	"rfd59/go-linky/cmd/go-linky/core/linky"
	"rfd59/go-linky/cmd/go-linky/infra"
	"rfd59/go-linky/cmd/go-linky/models"

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
		return nil, errors.New("No datasets found in the frame")
	}

	return mode.LoadTiC(ds), nil
}

func isValid(frame []byte) (bool, error) {
	if len(frame) == 0 {
		return false, errors.New("Frame is empty")
	}
	if frame[0] != 0x02 {
		return false, errors.New("Frame does not start with STX (0x02) character")
	}
	if frame[len(frame)-1] != 0x03 {
		return false, errors.New("Frame does not end with ETX (0x03) character")
	}
	return true, nil
}

func (s *LinkyService) OpenPort(settings *models.Serial, serial infra.ISerialInfra) (serial.Port, error) {
	// Open the serial port with the specified mode
	port, err := serial.Open(settings.Port, settings.Mode)
	if err != nil {
		return nil, fmt.Errorf("Failed to open the serial port %q: %w", settings.Port, err)
	}

	return port, nil
}
