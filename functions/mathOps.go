package functions

import (
	"math"
)

func Sub(a, b int) int {
	return a - b
}

func Add(a, b int) int {
	return a + b
}

func GetDifference(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}
