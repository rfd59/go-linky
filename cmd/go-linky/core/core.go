package core

import (
	"bufio"
	"fmt"
	"log/slog"
	"rfd59/go-linky/cmd/go-linky/core/linky"
	"rfd59/go-linky/cmd/go-linky/infra"
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/services"
)

func Run(settings *models.Settings, linkyService services.ILinkyService, mqttService services.IMqttService) error {
	slog.Debug("The serial port used is '" + settings.Linky.Serial.Port + "'.")

	// Open the serial port
	com, err := linkyService.OpenPort(&settings.Linky.Serial, &infra.SerialInfra{})
	if err != nil {
		return err
	}
	defer com.Close()

	// Load the buffer
	reader := bufio.NewReader(com)

	loop := 0
	for loop <= 5 {
		// Read until "End TeXt" ETX (0x03) character
		frame, err := reader.ReadBytes(0x03)
		if err != nil {
			return fmt.Errorf("port is unavailable to read from: %w", err)
		}

		// Check if the frame starts with "Start TeXt" STX (0x02) character to have a full frame
		if frame[0] == 0x02 {
			processing(frame, settings, linkyService, mqttService)
			return nil
		} else {
			// frame is incomplet. The next will be OK
			loop++
			slog.Debug("The frame is incomplete, reading the following.")
		}
	}

	return nil
}

func processing(frame []byte, settings *models.Settings, linkyService services.ILinkyService, mqttService services.IMqttService) {
	var mode linky.ILinkyMode

	slog.Debug("Processing the frame...")

	// Create the appropriate linky object based on the mode
	switch settings.Linky.Mode {
	case models.StandardMode:
		mode = &linky.Standard{}
	default:
		mode = &linky.Historic{}
	}

	tic, err := linkyService.ReadTic(frame, mode) // Analyze the frame in historic mode
	if err != nil {
		slog.Error("The TIC can't be read! [" + err.Error() + "]")
	} else {
		slog.Debug("The TIC is getted.", "tic", tic)
		slog.Debug("Publishing the TIC...")
		if err := mqttService.Publish(tic, &settings.Mqtt, mqttService.GetClient()); err != nil {
			slog.Error("The TIC can't be published! [" + err.Error() + "]")
		}
	}

	slog.Debug("The TIC has been published.")
}
