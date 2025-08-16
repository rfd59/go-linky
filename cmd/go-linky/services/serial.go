package services

import (
	"errors"
	"fmt"
	"log/slog"
	"rfd59/go-linky/cmd/go-linky/infra"
	"rfd59/go-linky/cmd/go-linky/models"
)

type ISerialService interface {
	DiscoverPort(settings *models.Serial, serial infra.ISerialInfra) error
}

type SerialService struct{}

func (s *SerialService) DiscoverPort(settings *models.Serial, serial infra.ISerialInfra) error {
	if settings.Port == "" {
		// Retrieve the port list
		ports, err := serial.GetPortsList()
		if err != nil {
			return fmt.Errorf("Failed to get the port list: %w", err)
		}
		// Check if any ports were found
		if len(ports) == 0 {
			return errors.New("No serial ports found")
		} else if len(ports) > 1 {
			slog.Warn("Multiple serial ports found, using the first one.")
		}

		settings.Port = ports[0]
	}

	return nil
}
