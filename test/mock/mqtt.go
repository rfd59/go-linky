package mock_test

import (
	"rfd59/go-linky/cmd/go-linky/models"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MqttClientMock is a mock implementation of the mqtt.Client interface
// used for testing purposes.
type MqttClientMock struct {
	isConnected      bool
	isConnectionOpen bool
	token            mqtt.Token
}

func InitMockMqttClient_IsConnected(mock bool) (m *MqttClientMock) {
	m = &MqttClientMock{}
	m.isConnected = mock

	return m
}

func InitMockMqttClient_Publish(mock mqtt.Token) (m *MqttClientMock) {
	m = &MqttClientMock{}
	m.isConnected = true
	m.token = mock

	return m
}

func (m *MqttClientMock) IsConnected() bool       { return m.isConnected }
func (m *MqttClientMock) IsConnectionOpen() bool  { return m.isConnectionOpen }
func (m *MqttClientMock) Connect() mqtt.Token     { return m.token }
func (m *MqttClientMock) Disconnect(quiesce uint) {}
func (m *MqttClientMock) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return m.token
}
func (m *MqttClientMock) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	return m.token
}
func (m *MqttClientMock) SubscribeMultiple(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token {
	return m.token
}
func (m *MqttClientMock) Unsubscribe(topics ...string) mqtt.Token             { return m.token }
func (m *MqttClientMock) AddRoute(topic string, callback mqtt.MessageHandler) {}
func (m *MqttClientMock) OptionsReader() mqtt.ClientOptionsReader             { return mqtt.ClientOptionsReader{} }

// MqttTokenMock is a mock implementation of the mqtt.Token interface
// used for testing purposes.
type MqttTokenMock struct {
	wait        bool
	waitTimeout bool
	err         error
}

func InitMockMqttToken(wait bool, err error) (m *MqttTokenMock) {
	m = &MqttTokenMock{}
	m.wait = wait
	m.waitTimeout = wait
	m.err = err

	return m
}

func (m *MqttTokenMock) Wait() bool                     { return m.wait }
func (m *MqttTokenMock) WaitTimeout(time.Duration) bool { return m.waitTimeout }
func (m *MqttTokenMock) Done() <-chan struct{}          { return nil }
func (m *MqttTokenMock) Error() error                   { return m.err }

// MqttServiceMock is a mock implementation of the services.IMqttService interface
// used for testing purposes.
type MqttServiceMock struct {
	topic string
	err   error
}

func InitMockMqttService(topic string, err error) (m *MqttServiceMock) {
	m = &MqttServiceMock{}
	m.topic = topic
	m.err = err

	return m
}

func (m *MqttServiceMock) GetClient() mqtt.Client                        { return nil }
func (m *MqttServiceMock) GetTopicName(topic string, adco string) string { return m.topic }
func (m *MqttServiceMock) Publish(tic *models.TiC, settings *models.MqttSettings, cli mqtt.Client) error {
	return m.err
}
