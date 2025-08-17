package infra

import "go.bug.st/serial"

type ISerialInfra interface {
	GetPortsList() ([]string, error)
	Open(portName string, mode *serial.Mode) (serial.Port, error)
}

type SerialInfra struct{}

func (i *SerialInfra) GetPortsList() ([]string, error) {
	return serial.GetPortsList()
}

func (i *SerialInfra) Open(portName string, mode *serial.Mode) (serial.Port, error) {
	return serial.Open(portName, mode)
}
