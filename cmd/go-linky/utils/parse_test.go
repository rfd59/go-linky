package utils_test

import (
	"rfd59/go-linky/cmd/go-linky/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_ParseUint8(t *testing.T) {
	assert := assert.New(t)

	type testCaseStruct struct {
		input    string
		expected uint8
	}

	for id, testCase := range map[string]testCaseStruct{
		"correct":    {input: "12", expected: 12},
		"invalid":    {input: "-34", expected: 0},
		"outOfRange": {input: "123456", expected: 0},
		"string":     {input: "azerty", expected: 0},
	} {
		t.Run(id, func(t *testing.T) {
			data := utils.ParseUint8(testCase.input)

			assert.Equal(testCase.expected, data)
		})
	}
}

func TestParse_ParseUint16(t *testing.T) {
	assert := assert.New(t)

	type testCaseStruct struct {
		input    string
		expected uint16
	}

	for id, testCase := range map[string]testCaseStruct{
		"correct":    {input: "1234", expected: 1234},
		"invalid":    {input: "-34", expected: 0},
		"outOfRange": {input: "123456", expected: 0},
		"string":     {input: "azerty", expected: 0},
	} {
		t.Run(id, func(t *testing.T) {
			data := utils.ParseUint16(testCase.input)

			assert.Equal(testCase.expected, data)
		})
	}
}

func TestParse_ParseUint32(t *testing.T) {
	assert := assert.New(t)

	type testCaseStruct struct {
		input    string
		expected uint32
	}

	for id, testCase := range map[string]testCaseStruct{
		"correct":    {input: "123456789", expected: 123456789},
		"invalid":    {input: "-34", expected: 0},
		"outOfRange": {input: "12345678900", expected: 0},
		"string":     {input: "azerty", expected: 0},
	} {
		t.Run(id, func(t *testing.T) {
			data := utils.ParseUint32(testCase.input)

			assert.Equal(testCase.expected, data)
		})
	}
}
