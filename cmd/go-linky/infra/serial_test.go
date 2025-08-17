package infra_test

import (
	"os"
	"rfd59/go-linky/cmd/go-linky/infra"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerial_GetPortsList(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Test the GetPortsList function
	service := &infra.SerialInfra{}
	ports, err := service.GetPortsList()

	// Assert the expected behavior
	require.NoError(err)
	require.NoError(err, "Expected no error when getting the serial port")
	if os.Getenv("CI_JOB_ID") != "" { // If running in GitLab Pipeline, we expect a specific number of ports
		assert.Len(ports, 4, "Expected four ports to be returned when running the test")
	} else {
		assert.Empty(ports, "Expected no ports to be returned when running the test")
	}
}

func TestSerial_Open(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Test the Open function
	service := &infra.SerialInfra{}
	ports, err := service.Open("COM1", nil)

	// Assert the expected behavior
	require.Error(err)
	assert.Nil(ports, "Expected no port to be returned when running the test")
}
