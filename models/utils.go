package models

import "math"

func half(width int) int {
	return int(math.Floor(float64(width) / 2.0))
}

func quarter(width int) int {
	return int(math.Floor(float64(width) / 4.0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
