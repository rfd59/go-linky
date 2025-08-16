package services_test

import (
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/services"
	mock_test "rfd59/go-linky/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMqtt_Publish_ConnectionFailed(t *testing.T) {
	require := require.New(t)

	// Test the GetSerialPort function
	settings := &models.MqttSettings{}
	service := services.NewMqttService(models.MqttSettings{})

	err := service.Publish(nil, settings, service.GetClient())

	// Assert the expected behavior
	require.Error(err)
	require.EqualError(err, "failed to etablish a connection to MQTT: no servers defined to connect to")
}

func TestMqtt_Publish_MessageFailed(t *testing.T) {
	require := require.New(t)

	// Mock
	mClient := mock_test.InitMockMqttClient_IsConnected(true)

	// Test the GetSerialPort function
	settings := &models.MqttSettings{}
	service := &services.MqttService{}

	err := service.Publish(nil, settings, mClient)

	// Assert the expected behavior
	require.Error(err)
	require.EqualError(err, "message can't be build: null value")
}

func TestMqtt_Publish_Failed(t *testing.T) {
	require := require.New(t)

	// Mock
	mToken := mock_test.InitMockMqttToken(true, mock_test.MockError)
	mClient := mock_test.InitMockMqttClient_Publish(mToken)

	// Test the GetSerialPort function
	settings := &models.MqttSettings{}
	service := &services.MqttService{}

	err := service.Publish(&models.TiC{}, settings, mClient)

	// Assert the expected behavior
	require.Error(err)
	require.EqualError(err, "failed to publish the message to MQTT: mock error...")
}

func TestMqtt_Publish_Success(t *testing.T) {
	require := require.New(t)

	// Mock
	mToken := mock_test.InitMockMqttToken(false, nil)
	mClient := mock_test.InitMockMqttClient_Publish(mToken)

	// Test the GetSerialPort function
	settings := &models.MqttSettings{}
	service := &services.MqttService{}

	err := service.Publish(&models.TiC{}, settings, mClient)

	// Assert the expected behavior
	require.NoError(err)
}

func TestMqtt_GetTopicName(t *testing.T) {
	assert := assert.New(t)

	type testCaseStruct struct {
		topic    string
		adco     string
		expected string
	}

	for id, testCase := range map[string]testCaseStruct{
		"defaut": {expected: "linky/default"},
		"adco":   {adco: "123456789", expected: "linky/123456789"},
		"topic":  {topic: "my/custom", expected: "my/custom"},
		"full":   {topic: "my/custom", adco: "123456789", expected: "my/custom"},
	} {
		t.Run(id, func(t *testing.T) {
			// Test the GetTopicName function
			service := &services.MqttService{}
			data := service.GetTopicName(testCase.topic, testCase.adco)

			// Assert the expected behavior
			assert.NotNil(data)
			assert.Equal(testCase.expected, data)
		})
	}
}

func TestMqtt_MqttClient(t *testing.T) {
	assert := assert.New(t)

	// Test the GetPortsList function
	service := services.NewMqttService(models.MqttSettings{})

	// Assert the expected behavior
	assert.NotNil(service.GetClient())
	assert.False(service.GetClient().IsConnected())
}
