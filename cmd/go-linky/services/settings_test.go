package services_test

import (
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/services"
	mock_test "rfd59/go-linky/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.bug.st/serial"
)

func TestSettings_NewSettingsService_Default(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Mock the SerialService
	mSerialService := mock_test.InitMockSerialService(nil)

	// Test the NewSettingsService function
	service, err := services.NewSettingsService(mSerialService)

	/// Assert the expected behavior
	require.NoError(err)
	s := service.Get() // get the settings loaded
	assert.Equal(models.LinkySettings{Mode: models.HistoricMode, Frequency: 10, Serial: models.Serial{Mode: &serial.Mode{BaudRate: 1200, DataBits: 7, StopBits: serial.OneStopBit, Parity: serial.EvenParity}}}, s.Linky)
	assert.Equal(models.MqttSettings{Protocol: "tcp", Host: "localhost", Port: 1883}, s.Mqtt)
}

func TestSettings_NewSettingsService_EnvVar(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Mock the SerialService
	mSerialService := mock_test.InitMockSerialService(nil)

	// Set Environments variables
	t.Setenv("GOLINKY_LINKY_SERIAL_PORT", "COM99")
	t.Setenv("GOLINKY_LINKY_MODE", "STANDARD")
	t.Setenv("GOLINKY_LINKY_FREQUENCY", "300")
	t.Setenv("GOLINKY_MQTT_PROTOCOL", "tcps")
	t.Setenv("GOLINKY_MQTT_HOST", "mqtt.domain.local")
	t.Setenv("GOLINKY_MQTT_PORT", "1884")
	t.Setenv("GOLINKY_MQTT_USERNAME", "mqtt")
	t.Setenv("GOLINKY_MQTT_PASSWORD", "secret")
	t.Setenv("GOLINKY_MQTT_TOPIC", "my/topic")

	// Test the NewSettingsService function
	service, err := services.NewSettingsService(mSerialService)

	/// Assert the expected behavior
	require.NoError(err)
	s := service.Get() // get the settings loaded
	assert.Equal(models.LinkySettings{Mode: models.StandardMode, Frequency: 300, Serial: models.Serial{Port: "COM99", Mode: &serial.Mode{BaudRate: 9600, DataBits: 7, StopBits: serial.OneStopBit, Parity: serial.EvenParity}}}, s.Linky)
	assert.Equal(models.MqttSettings{Protocol: "tcps", Host: "mqtt.domain.local", Port: 1884, Username: "mqtt", Password: "secret", Topic: "my/topic"}, s.Mqtt)
}

func TestSettings_NewSettingsService_Error(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Mock the SerialService
	mSerialService := mock_test.InitMockSerialService(mock_test.MockError)

	// Test the NewSettingsService function
	service, err := services.NewSettingsService(mSerialService)

	/// Assert the expected behavior
	require.Error(err)
	require.EqualError(err, "discover serial port failed: mock error...")
	assert.Nil(service)
}
