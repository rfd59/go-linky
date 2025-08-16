package services

import (
	"fmt"
	"os"
	"rfd59/go-linky/cmd/go-linky/infra"
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/utils"
	"strings"

	"go.bug.st/serial"
)

type ISettingsService interface {
	Get() *models.Settings
}

type SettingsService struct {
	settings *models.Settings
}

func NewSettingsService(serialService ISerialService) (s *SettingsService, err error) {
	s = &SettingsService{
		settings: &models.Settings{},
	}

	s.loadSettings()
	if err := serialService.DiscoverPort(&s.settings.Linky.Serial, &infra.SerialInfra{}); err != nil {
		return nil, fmt.Errorf("Discover serial port failed: %w", err)
	}

	return s, nil
}

func (s *SettingsService) loadSettings() {
	s.loadLinkySettings()
	s.loadMqttSettings()
}

func (s *SettingsService) loadLinkySettings() {
	s.getLinkyModeSetting()
	s.settings.Linky.Frequency = utils.ParseUint16(s.getEnvironmentSetting("GOLINKY_LINKY_FREQUENCY", "10"))
	s.getLinkySerialSetting()
}

func (s *SettingsService) loadMqttSettings() {
	s.settings.Mqtt.Protocol = s.getEnvironmentSetting("GOLINKY_MQTT_PROTOCOL", "tcp")
	s.settings.Mqtt.Host = s.getEnvironmentSetting("GOLINKY_MQTT_HOST", "localhost")
	s.settings.Mqtt.Port = utils.ParseUint16(s.getEnvironmentSetting("GOLINKY_MQTT_PORT", "1883"))
	s.settings.Mqtt.Username = s.getEnvironmentSetting("GOLINKY_MQTT_USERNAME", "")
	s.settings.Mqtt.Password = s.getEnvironmentSetting("GOLINKY_MQTT_PASSWORD", "")
	s.settings.Mqtt.Topic = s.getEnvironmentSetting("GOLINKY_MQTT_TOPIC", "")
}

func (s *SettingsService) getEnvironmentSetting(name string, def string) string {
	if os.Getenv(name) != "" {
		return os.Getenv(name)
	} else {
		return def
	}
}

func (s *SettingsService) getLinkyModeSetting() {
	if strings.ToUpper(s.getEnvironmentSetting("GOLINKY_LINKY_MODE", "HISTORIQUE")) == "STANDARD" {
		s.settings.Linky.Mode = models.StandardMode
	} else {
		s.settings.Linky.Mode = models.HistoricMode
	}
}

func (s *SettingsService) getLinkySerialSetting() {
	s.settings.Linky.Serial.Port = s.getEnvironmentSetting("GOLINKY_LINKY_SERIAL_PORT", "")

	switch s.settings.Linky.Mode {
	case models.HistoricMode:
		s.settings.Linky.Serial.Mode = &serial.Mode{
			BaudRate: 1200,
			Parity:   serial.EvenParity,
			DataBits: 7,
			StopBits: serial.OneStopBit,
		}
	case models.StandardMode:
		s.settings.Linky.Serial.Mode = &serial.Mode{
			BaudRate: 9600,
			Parity:   serial.EvenParity,
			DataBits: 7,
			StopBits: serial.OneStopBit,
		}
	}
}

func (s *SettingsService) Get() *models.Settings {
	return s.settings
}
