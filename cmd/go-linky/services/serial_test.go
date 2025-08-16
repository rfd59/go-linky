package services_test

import (
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/services"
	mock_test "rfd59/go-linky/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerial_DiscoverPort_Ports(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	for id, testCase := range map[string][]string{
		"simple":   {"COM1"},
		"multiple": {"/dev/serial0", "/dev/serial1"},
	} {
		t.Run(id, func(t *testing.T) {
			// Mock the SerialInfra
			mInfra := mock_test.InitMockSerialInfra_GetPortsList(mock_test.SerialInfra_GetPortsList{Ports: testCase, Error: nil})

			// Test the GetSerialPort function
			settings := &models.Serial{Port: "", Mode: nil}
			service := &services.SerialService{}

			err := service.DiscoverPort(settings, mInfra)

			// Assert the expected behavior
			require.NoError(err, "Expected no error when getting the serial port")
			assert.Equal(testCase[0], settings.Port, "Expected the first port to be returned")
		})
	}
}

func TestSerial_GetSerialPort_Empty(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Mock the SerialInfra
	mInfra := mock_test.InitMockSerialInfra_GetPortsList(mock_test.SerialInfra_GetPortsList{Ports: []string{}, Error: nil})

	// Test the GetSerialPort function
	settings := &models.Serial{Port: "", Mode: nil}
	service := &services.SerialService{}

	err := service.DiscoverPort(settings, mInfra)

	// Assert the expected behavior
	require.Error(err, "Expected an error when no ports are available")
	require.EqualError(err, "no serial ports found", "Expected specific error message when no ports are available")
	assert.Empty(settings.Port, "Expected no port to be returned when no ports are available")
}

func TestSerial_GetSerialPort_Error(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Mock the SerialInfra
	mInfra := mock_test.InitMockSerialInfra_GetPortsList(mock_test.SerialInfra_GetPortsList{Ports: nil, Error: mock_test.ErrMockor})

	// Test the GetSerialPort function
	settings := &models.Serial{Port: "", Mode: nil}
	service := &services.SerialService{}

	err := service.DiscoverPort(settings, mInfra)

	// Assert the expected behavior
	require.Error(err, "Expected an error when no ports are available")
	require.EqualError(err, "failed to get the port list: mock error...", "Expected specific error message when GetPortsList fails")
	assert.Empty(settings.Port, "Expected no port to be returned when no ports are available")
}
