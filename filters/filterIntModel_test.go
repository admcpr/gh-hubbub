package filters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFilterIntModel(t *testing.T) {
	tests := []struct {
		tab   string
		title string
		from  int
		to    int
	}{
		{"Tab", "Test 1", 0, 10},
		{"Tab", "Test 2", -5, 5},
		{"Tab", "Test 3", 100, 200},
		{"Tab", "Test 4", -100, 0},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			m := NewIntModel(tt.tab, tt.title, tt.from, tt.to)

			assert.Equal(t, m.Tab, tt.tab)
			assert.Equal(t, m.Title, tt.title)
			assert.Equal(t, m.fromInput.Placeholder, fmt.Sprint(tt.from))
			assert.Equal(t, m.toInput.Placeholder, fmt.Sprint(tt.to))
		})
	}
}

func TestIntValidator(t *testing.T) {
	tests := []struct {
		input  string
		prompt string
		want   error
	}{
		{"123", "Test 1", nil},
		{"-5", "Test 2", nil},
		{"abc", "Test 3", fmt.Errorf("please enter an integer for the `Test 3` value")},
		{"1.23", "Test 4", fmt.Errorf("please enter an integer for the `Test 4` value")},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s: %s", tt.prompt, tt.input), func(t *testing.T) {
			got := intValidator(tt.input, tt.prompt)

			assert.Equal(t, got, tt.want, fmt.Sprintf("got %q, want %q", got, tt.want))
		})
	}
}
