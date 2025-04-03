package utils

import (
	"math"
)

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Numeric interface {
	Integer | Float
}

// Abs returns the absolute value of x.
func Abs[T Numeric](x T) T {
	var zero T
	if x < zero {
		return -x
	}
	return x
}

// Round returns the nearest integer, rounding half away from zero.
func Round[T Integer](x float64) T {
	return T(math.Round(x))
}

// RoundUp returns the nearest large integer (ceil).
func RoundUp[T Integer](x float64) T {
	return T(math.Ceil(x))
}

// RoundDown returns the nearest small integer.
func RoundDown[T Integer](x float64) T {
	return T(x)
}

// Min returns the smallest value of numbers or zero.
func Min[T Numeric](numbers ...T) T {
	if len(numbers) == 0 {
		var zero T
		return zero
	}

	res := numbers[0]
	for _, num := range numbers {
		if num < res {
			res = num
		}
	}
	return res
}

// Max returns the largest value of numbers or zero.
func Max[T Numeric](numbers ...T) T {
	if len(numbers) == 0 {
		var zero T
		return zero
	}

	res := numbers[0]
	for _, num := range numbers {
		if num > res {
			res = num
		}
	}
	return res
}
