package mock_test

import (
	"rfd59/go-linky/cmd/go-linky/core/linky"
	"rfd59/go-linky/cmd/go-linky/infra"
	"rfd59/go-linky/cmd/go-linky/models"

	"go.bug.st/serial"
)

// SerialInfraMock is a mock implementation of the utils.ISerialTool interface
// used for testing purposes.
type LinkyModeMock struct {
	loadDatasets []models.LinkyDataset
	loadTiC      *models.TiC
}

func InitMockLinkyMode(loadDatasets []models.LinkyDataset, loadTiC models.TiC) (m *LinkyModeMock) {
	m = &LinkyModeMock{}
	m.loadDatasets = loadDatasets
	m.loadTiC = &loadTiC

	return m
}

func (m *LinkyModeMock) LoadDatasets(frame string) []models.LinkyDataset {
	return m.loadDatasets
}

func (m *LinkyModeMock) LoadTiC(ds []models.LinkyDataset) *models.TiC {
	return m.loadTiC
}

// LinkyServiceMock is a mock implementation of the services.ILinkyService interface
// used for testing purposes.
type LinkyServiceMock struct {
	openPort LinkyService_OpenPort
	readTiC  LinkyService_ReadTiC
}

type LinkyService_OpenPort struct {
	Port serial.Port
	Err  error
}

type LinkyService_ReadTiC struct {
	TiC models.TiC
	Err error
}

func InitMockLinkyService(openPort LinkyService_OpenPort, readTiC LinkyService_ReadTiC) (m *LinkyServiceMock) {
	m = &LinkyServiceMock{}
	m.openPort = openPort
	m.readTiC = readTiC

	return m
}

func (m *LinkyServiceMock) OpenPort(settings *models.Serial, serial infra.ISerialInfra) (serial.Port, error) {
	return m.openPort.Port, m.openPort.Err
}

func (m *LinkyServiceMock) ReadTic(frame []byte, mode linky.ILinkyMode) (*models.TiC, error) {
	return &m.readTiC.TiC, m.readTiC.Err
}
