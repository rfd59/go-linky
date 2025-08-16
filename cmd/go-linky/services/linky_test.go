package services_test

import (
	"fmt"
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/services"
	mock_test "rfd59/go-linky/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.bug.st/serial"
)

func TestLinky_ReadTiC_FrameNotValid(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	type testCaseStruct struct {
		frame    []byte
		expected string
	}

	for id, testCase := range map[string]testCaseStruct{
		"empty":  {frame: []byte{}, expected: "frame is empty"},
		"no STX": {frame: []byte{10, 56, 13, 03}, expected: "frame does not start with STX (0x02) character"},
		"no ETX": {frame: []byte{02, 10, 56, 13}, expected: "frame does not end with ETX (0x03) character"},
	} {
		t.Run(id, func(t *testing.T) {
			// Test the ReadTic function
			service := &services.LinkyService{}
			data, err := service.ReadTic(testCase.frame, nil)

			// Assert the expected behavior
			require.Error(err, "Expected an error for invalid frame")
			require.EqualError(err, testCase.expected, "Error message does not match expected")
			assert.Nil(data, "Expected data to be nil for invalid frame")
		})
	}
}

func TestLinky_ReadTiC_Error(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Mock the LinkyMode
	mLinkyMode := mock_test.InitMockLinkyMode([]models.LinkyDataset{}, models.TiC{})

	// Test the ReadTic function
	service := &services.LinkyService{}
	data, err := service.ReadTic([]byte{2, 10, 65, 68, 67, 32, 79, 13, 3}, mLinkyMode)

	// Assert the expected behavior
	require.Error(err, "Expected an error for invalid dataset format")
	require.EqualError(err, "no datasets found in the frame", "Error message does not match expected")
	assert.Nil(data, "Expected data to be nil for invalid frame")
}

func TestLinky_ReadTic(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Mock the LinkyMode
	mLinkyMode := mock_test.InitMockLinkyMode([]models.LinkyDataset{{Label: "UnitTest", Data: "Mock"}}, models.TiC{})

	// Test the ReadTic function
	service := &services.LinkyService{}
	data, err := service.ReadTic([]byte{2, 10, 65, 68, 67, 32, 79, 13, 3}, mLinkyMode)

	// Assert the expected behavior
	require.NoError(err, "Expected no error")
	assert.NotNil(data, "Expected data to be not nil")
}

func TestLinky_OpenPort_Open(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Mock the SerialInfra
	mInfra := mock_test.InitMockSerialInfra_Open(map[string]mock_test.SerialInfra_Open{"COM1": {
		Port:  &mock_test.SerialPortMock{}, // Mocked port, can be nil for testing
		Error: nil,
	}})

	// Test the OpenSerialPort function
	settings := &models.Serial{Port: "COM1", Mode: &serial.Mode{}}
	service := &services.LinkyService{}

	port, err := service.OpenPort(settings, mInfra)

	// Assert the expected behavior
	require.NoError(err, "Expected no error when opening the serial port")
	assert.NotNil(port, "Expected a valid port to be returned")
}

func TestLinky_OpenPort_Error(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Mock the SerialInfra
	mInfra := mock_test.InitMockSerialInfra_Open(map[string]mock_test.SerialInfra_Open{"COM1": {
		Port:  nil,
		Error: mock_test.MockError,
	}})

	// Test the OpenSerialPort function
	settings := &models.Serial{Port: "COM1", Mode: nil}
	service := &services.LinkyService{}

	port, err := service.OpenPort(settings, mInfra)

	// Assert the expected behavior
	require.Error(err, "Expected error when opening the serial port")
	require.EqualError(err, fmt.Sprintf("failed to open the serial port %q: %s", "COM1", "mock error..."), "Expected specific error message when opening the serial port fails")
	assert.Nil(port, "Expected no port to be returned when opening the serial port fails")
}
