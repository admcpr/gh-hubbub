package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFilterBool(t *testing.T) {
	type args struct {
		tab   string
		name  string
		value bool
	}
	tests := []struct {
		args args
		text string
	}{
		{args: args{tab: "tab", name: "name", value: true}, text: "tab > name = Yes"},
		{args: args{tab: "tab", name: "name", value: false}, text: "tab > name = No"},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {

			got := NewFilterBool(tt.args.tab, tt.args.name, tt.args.value)
			assert.Equal(t, tt.args.tab, got.GetTab())
			assert.Equal(t, tt.args.name, got.GetName())
			assert.Equal(t, tt.text, got.String())
		})
	}
}
