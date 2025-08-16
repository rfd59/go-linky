package core_test

import (
	"errors"
	"log/slog"
	"rfd59/go-linky/cmd/go-linky/core"
	"rfd59/go-linky/cmd/go-linky/models"
	"rfd59/go-linky/cmd/go-linky/services"
	mock_test "rfd59/go-linky/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thejerf/slogassert"
)

func TestCore_Run_OpenPortFailed(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Test the GetPortsList function
	settings := &models.Settings{Linky: models.LinkySettings{Serial: models.Serial{Port: "COM1"}}}
	linky := &services.LinkyService{}
	mqtt := &services.MqttService{}

	err := core.Run(settings, linky, mqtt)

	// Assert the expected behavior
	require.Error(err)
	assert.EqualError(err, "failed to open the serial port \"COM1\": no such file or directory")
}

func TestCore_Run_ReadPortFailed(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Test the GetPortsList function
	settings := &models.Settings{}
	linky := mock_test.InitMockLinkyService(mock_test.LinkyService_OpenPort{Port: &mock_test.SerialPortMock{}}, mock_test.LinkyService_ReadTiC{})
	mqtt := &services.MqttService{}

	err := core.Run(settings, linky, mqtt)

	// Assert the expected behavior
	require.Error(err)
	assert.EqualError(err, "port is unavailable to read from: multiple Read calls return no data or error")
}

func TestCore_Run_FrameLoop(t *testing.T) {
	require := require.New(t)

	// Test the GetPortsList function
	settings := &models.Settings{}
	port := mock_test.InitMockSerialPort_Read([]byte{32, 65, 32, 44, 13, 3}, nil)
	linky := mock_test.InitMockLinkyService(mock_test.LinkyService_OpenPort{Port: port}, mock_test.LinkyService_ReadTiC{})
	mqtt := &services.MqttService{}

	err := core.Run(settings, linky, mqtt)

	// Assert the expected behavior
	require.NoError(err)
}

func TestCore_Run_ProcessingTicFailed(t *testing.T) {
	require := require.New(t)

	// Testing handler for slog.
	log := slogassert.New(t, slog.LevelError, nil)
	logger := slog.New(log)
	slog.SetDefault(logger)

	// Test the GetPortsList function
	settings := &models.Settings{}
	port := mock_test.InitMockSerialPort_Read([]byte{2, 10, 72, 72, 80, 72, 67, 32, 65, 32, 44, 13, 3}, nil)
	linky := mock_test.InitMockLinkyService(mock_test.LinkyService_OpenPort{Port: port}, mock_test.LinkyService_ReadTiC{TiC: models.TiC{}, Err: errors.New("mock error...")})
	mqtt := &services.MqttService{}

	err := core.Run(settings, linky, mqtt)

	// Assert the expected behavior
	require.NoError(err)
	log.AssertPrecise(slogassert.LogMessageMatch{Message: "The TIC can't be read! [mock error...]", Level: slog.LevelError})
}

func TestCore_Run_ProcessingPublishFailed(t *testing.T) {
	require := require.New(t)

	// Testing handler for slog.
	log := slogassert.New(t, slog.LevelError, nil)
	logger := slog.New(log)
	slog.SetDefault(logger)

	// Test the GetPortsList function
	settings := &models.Settings{Linky: models.LinkySettings{Mode: models.StandardMode}}
	port := mock_test.InitMockSerialPort_Read([]byte{2, 10, 72, 72, 80, 72, 67, 32, 65, 32, 44, 13, 3}, nil)
	linky := mock_test.InitMockLinkyService(mock_test.LinkyService_OpenPort{Port: port}, mock_test.LinkyService_ReadTiC{TiC: models.TiC{}})
	mqtt := mock_test.InitMockMqttService("my/topic", errors.New("mock error..."))

	err := core.Run(settings, linky, mqtt)

	// Assert the expected behavior
	require.NoError(err)
	log.AssertPrecise(slogassert.LogMessageMatch{Message: "The TIC can't be published! [mock error...]", Level: slog.LevelError})
}
