package utils_test

import (
	"testing"

	"github.com/go-universal/utils"
	"github.com/stretchr/testify/assert"
)

func TestPointerOf(t *testing.T) {
	value := 42
	ptr := utils.PointerOf(value)
	assert.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)
}

func TestSafeValue(t *testing.T) {
	var value *int
	assert.Equal(t, 0, utils.SafeValue(value))

	val := 42
	value = &val
	assert.Equal(t, val, utils.SafeValue(value))
}

func TestValueOf(t *testing.T) {
	var value *int
	fallback := 42
	assert.Equal(t, fallback, utils.ValueOf(value, fallback))

	val := 24
	value = &val
	assert.Equal(t, val, utils.ValueOf(value, fallback))
}

func TestAlter(t *testing.T) {
	var value *int
	fallback := 42
	assert.Equal(t, fallback, utils.Alter(value, fallback))

	val := 0
	value = &val
	assert.Equal(t, fallback, utils.Alter(value, fallback))

	val = 24
	value = &val
	assert.Equal(t, val, utils.Alter(value, fallback))
}

func TestNullableOf(t *testing.T) {
	var value *int
	assert.Nil(t, utils.NullableOf(value))

	val := 0
	value = &val
	assert.Nil(t, utils.NullableOf(value))

	val = 24
	value = &val
	assert.Equal(t, value, utils.NullableOf(value))
}

func TestIsEmpty(t *testing.T) {
	var value *int
	assert.True(t, utils.IsEmpty(value))

	val := 0
	value = &val
	assert.True(t, utils.IsEmpty(value))

	val = 24
	value = &val
	assert.False(t, utils.IsEmpty(value))
}

func TestIsSame(t *testing.T) {
	var a, b *int
	assert.True(t, utils.IsSame(a, b))

	valA := 42
	a = &valA
	assert.False(t, utils.IsSame(a, b))

	valB := 42
	b = &valB
	assert.True(t, utils.IsSame(a, b))

	valB = 24
	assert.False(t, utils.IsSame(a, b))
}
