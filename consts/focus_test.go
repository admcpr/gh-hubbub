package consts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFocus_Next(t *testing.T) {
	tests := []struct {
		input    Focus
		expected Focus
	}{
		{input: FocusList, expected: FocusTabs},
		{input: FocusTabs, expected: FocusFilter},
		{input: FocusFilter, expected: FocusFilter},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Calling Next() from %d should return %d", tt.input, tt.expected), func(t *testing.T) {
			got := tt.input.Next()
			assert.Equal(t, got, tt.expected)
		})
	}
}

func TestFocus_Prev(t *testing.T) {
	tests := []struct {
		input    Focus
		expected Focus
	}{
		{input: FocusList, expected: FocusList},
		{input: FocusTabs, expected: FocusList},
		{input: FocusFilter, expected: FocusTabs},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Calling Prev() from %d should return %d", tt.input, tt.expected), func(t *testing.T) {
			got := tt.input.Prev()
			assert.Equal(t, got, tt.expected)
		})
	}
}
