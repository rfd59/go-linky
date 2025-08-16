package services

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/utils"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type IMqttService interface {
	GetClient() mqtt.Client
	GetTopicName(topic string, adco string) string
	Publish(tic *models.TiC, settings *models.MqttSettings, cli mqtt.Client) error
}

type MqttService struct {
	cli mqtt.Client
}

func NewMqttService(settings models.MqttSettings) *MqttService {
	opt := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("%s://%s:%d", settings.Protocol, settings.Host, settings.Port)).
		SetUsername(settings.Username).
		SetPassword(settings.Password).
		SetClientID("go-linky")

	return &MqttService{
		cli: mqtt.NewClient(opt),
	}
}

func (s *MqttService) Publish(tic *models.TiC, settings *models.MqttSettings, cli mqtt.Client) error {
	if !cli.IsConnected() {
		slog.Debug("Not connected! Connecting...")
		// Connection to MQTT broker
		if token := cli.Connect(); token.Wait() && token.Error() != nil {
			return fmt.Errorf("failed to etablish a connection to MQTT: %w", token.Error())
		}
	}

	// Convert the TIC to JSON object
	msg, err := json.Marshal(tic)
	if err != nil || string(msg) == "null" {
		if err == nil {
			err = utils.ErrNullValue
		}
		return fmt.Errorf("message can't be build: %w", err)
	}

	// Publish
	if token := cli.Publish(s.GetTopicName(settings.Topic, tic.ADCO), 0, false, string(msg)); token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to publish the message to MQTT: %w", token.Error())
	}

	return nil
}

func (s *MqttService) GetTopicName(topic string, adco string) string {
	if topic == "" {
		if adco == "" {
			return "linky/default"
		} else {
			return "linky/" + adco
		}
	} else {
		return topic
	}
}

func (s *MqttService) GetClient() mqtt.Client {
	return s.cli
}
