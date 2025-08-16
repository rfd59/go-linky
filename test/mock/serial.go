package mock_test

import (
	"rfd59/go-linky/cmd/go-linky/infra"
	"rfd59/go-linky/cmd/go-linky/models"
	"time"

	"go.bug.st/serial"
)

// SerialInfraMock is a mock implementation of the utils.ISerialTool interface
// used for testing purposes.
type SerialInfraMock struct {
	getPortsList SerialInfra_GetPortsList
	open         map[string]SerialInfra_Open
}

type SerialInfra_GetPortsList struct {
	Ports []string
	Error error
}

type SerialInfra_Open struct {
	Port  serial.Port
	Error error
}

func InitMockSerialInfra_GetPortsList(mock SerialInfra_GetPortsList) (m *SerialInfraMock) {
	m = &SerialInfraMock{}
	m.getPortsList = mock

	return m
}

func (m *SerialInfraMock) GetPortsList() ([]string, error) {
	return m.getPortsList.Ports, m.getPortsList.Error
}

func InitMockSerialInfra_Open(mock map[string]SerialInfra_Open) (m *SerialInfraMock) {
	m = &SerialInfraMock{}
	m.open = mock

	return m
}

func (m *SerialInfraMock) Open(portName string, mode *serial.Mode) (serial.Port, error) {
	return m.open[portName].Port, m.open[portName].Error
}

// SerialPortMock is a mock implementation of the serial.Port interface
// used for testing purposes.
type SerialPortMock struct {
	readData []byte
	readErr  error
}

func InitMockSerialPort_Read(data []byte, err error) (m *SerialPortMock) {
	m = &SerialPortMock{}
	m.readData = data
	m.readErr = err

	return m
}

func (m *SerialPortMock) SetMode(mode *serial.Mode) error {
	return nil
}
func (m *SerialPortMock) Read(p []byte) (n int, err error) {
	return copy(p, m.readData), m.readErr
}
func (m *SerialPortMock) Write(p []byte) (n int, err error) {
	return len(p), nil
}
func (m *SerialPortMock) Drain() error {
	return nil
}
func (m *SerialPortMock) ResetInputBuffer() error {
	return nil
}
func (m *SerialPortMock) ResetOutputBuffer() error {
	return nil
}
func (m *SerialPortMock) SetDTR(dtr bool) error {
	return nil
}
func (m *SerialPortMock) SetRTS(rts bool) error {
	return nil
}
func (m *SerialPortMock) GetModemStatusBits() (*serial.ModemStatusBits, error) {
	return &serial.ModemStatusBits{}, nil
}
func (m *SerialPortMock) SetReadTimeout(t time.Duration) error {
	return nil
}
func (m *SerialPortMock) Close() error {
	return nil
}
func (m *SerialPortMock) Break(d time.Duration) error {
	return nil
}

// SerialServiceMock is a mock implementation of the services.ISerialService interface
// used for testing purposes.
type SerialServiceMock struct {
	port serial.Port
	err  error
}

func InitMockSerialService(err error) (m *SerialServiceMock) {
	m = &SerialServiceMock{}
	m.err = err

	return m
}

func (m *SerialServiceMock) DiscoverPort(settings *models.Serial, serial infra.ISerialInfra) error {
	return m.err
}
