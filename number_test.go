package utils_test

import (
	"testing"

	"github.com/go-universal/utils"
	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{input: int(-5), expected: int(5)},
		{input: int(5), expected: int(5)},
		{input: float64(-5.5), expected: float64(5.5)},
		{input: float64(5.5), expected: float64(5.5)},
	}

	for _, test := range tests {
		switch v := test.input.(type) {
		case int:
			assert.Equal(t, test.expected, utils.Abs(v), "Abs(%v)", v)
		case float64:
			assert.Equal(t, test.expected, utils.Abs(v), "Abs(%v)", v)
		}
	}
}

func TestRoundUp(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{input: 5.1, expected: 6},
		{input: 5.9, expected: 6},
		{input: -5.1, expected: -5},
		{input: -5.9, expected: -5},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, utils.RoundUp[int](test.input), "RoundUp(%v)", test.input)
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{input: 5.5, expected: 6},
		{input: 5.4, expected: 5},
		{input: -5.5, expected: -6},
		{input: -5.4, expected: -5},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, utils.Round[int](test.input), "Round(%v)", test.input)
	}
}

func TestRoundDown(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{input: 5.9, expected: 5},
		{input: 5.1, expected: 5},
		{input: -5.9, expected: -5},
		{input: -5.1, expected: -5},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, utils.RoundDown[int](test.input), "RoundDown(%v)", test.input)
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected interface{}
	}{
		{input: []interface{}{int(5), int(3), int(9)}, expected: int(3)},
		{input: []interface{}{float64(5.5), float64(3.3), float64(9.9)}, expected: float64(3.3)},
	}

	for _, test := range tests {
		switch test.input[0].(type) {
		case int:
			assert.Equal(t, test.expected, utils.Min(test.input[0].(int), test.input[1].(int), test.input[2].(int)), "Min(%v)", test.input)
		case float64:
			assert.Equal(t, test.expected, utils.Min(test.input[0].(float64), test.input[1].(float64), test.input[2].(float64)), "Min(%v)", test.input)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected interface{}
	}{
		{input: []interface{}{int(5), int(3), int(9)}, expected: int(9)},
		{input: []interface{}{float64(5.5), float64(3.3), float64(9.9)}, expected: float64(9.9)},
	}

	for _, test := range tests {
		switch test.input[0].(type) {
		case int:
			assert.Equal(t, test.expected, utils.Max(test.input[0].(int), test.input[1].(int), test.input[2].(int)), "Max(%v)", test.input)
		case float64:
			assert.Equal(t, test.expected, utils.Max(test.input[0].(float64), test.input[1].(float64), test.input[2].(float64)), "Max(%v)", test.input)
		}
	}
}
