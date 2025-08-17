package utils

import (
	"log/slog"
	"strconv"
)

func ParseUint8(s string) uint8 {
	return uint8(parseUint(s, 8))
}

func ParseUint16(s string) uint16 {
	return uint16(parseUint(s, 16))
}

func ParseUint32(s string) uint32 {
	return uint32(parseUint(s, 32))
}

func parseUint(s string, bitsize int) uint64 {
	if val, err := strconv.ParseUint(s, 10, bitsize); err == nil {
		return val
	} else {
		slog.Warn("Failed to parse value '"+s+"' to integer! ["+err.Error()+"]", s, err)
		return 0
	}
}
