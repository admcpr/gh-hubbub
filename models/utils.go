package models

func half(width int) int {
	return width / 2
}

func quarter(width int) int {
	return width / 4
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
