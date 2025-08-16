package models

import "go.bug.st/serial"

type Settings struct {
	Linky LinkySettings
	Mqtt  MqttSettings
}

type LinkySettings struct {
	Mode      Mode
	Frequency uint16
	Serial    Serial
}

type Mode uint8

const (
	HistoricMode = iota
	StandardMode
)

type Serial struct {
	Port string
	Mode *serial.Mode
}

type MqttSettings struct {
	Protocol string
	Host     string
	Port     uint16
	Username string
	Password string
	Topic    string
}
