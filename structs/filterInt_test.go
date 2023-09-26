package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFilterInt(t *testing.T) {
	type args struct {
		tab  string
		name string
		from int
		to   int
	}
	tests := []struct {
		args args
		text string
	}{
		{args: args{tab: "tab", name: "name", from: 0, to: 0}, text: "tab > name between 0 and 0"},
		{args: args{tab: "tab", name: "name", from: 0, to: 1}, text: "tab > name between 0 and 1"},
		{args: args{tab: "tab", name: "name", from: 1, to: 0}, text: "tab > name between 1 and 0"},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			got := NewFilterInt(tt.args.tab, tt.args.name, tt.args.from, tt.args.to)
			assert.Equal(t, tt.args.tab, got.GetTab())
			assert.Equal(t, tt.args.name, got.GetName())
			assert.Equal(t, tt.text, got.String())
		})
	}
}
