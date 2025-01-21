package shared

import "math"

func Half(width int) int {
	return int(math.Floor(float64(width) / 2.0))
}

func Quarter(width int) int {
	return int(math.Floor(float64(width) / 4.0))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
