package umath

import "math"

// Not quite efficient
func Min3(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}

func MinUint64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
