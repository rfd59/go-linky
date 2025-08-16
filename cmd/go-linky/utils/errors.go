package utils

import "errors"

var ErrNoDatasets = errors.New("no datasets found in the frame")
var ErrNoSerialPort = errors.New("no serial ports found")
var ErrInvalidDataset = errors.New("invalid dataset format")
var ErrInvalidChecksum = errors.New("invalid checksum length")
var ErrEmptyFrame = errors.New("frame is empty")
var ErrSTXFrame = errors.New("frame does not start with STX (0x02) character")
var ErrETXFrame = errors.New("frame does not end with ETX (0x03) character")
var ErrNullValue = errors.New("null value")
