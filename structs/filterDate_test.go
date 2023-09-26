package structs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFilterDate(t *testing.T) {
	type args struct {
		tab  string
		name string
		from string
		to   string
	}
	tests := []struct {
		args args
		text string
	}{
		{args: args{tab: "tab", name: "name", from: "2000-01-01", to: "2000-01-02"}, text: "tab > name between 2000-01-01 and 2000-01-02"},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			from, _ := time.Parse("2006-01-02", tt.args.from)
			to, _ := time.Parse("2006-01-02", tt.args.to)

			got := NewFilterDate(tt.args.tab, tt.args.name, from, to)
			assert.Equal(t, tt.args.tab, got.GetTab())
			assert.Equal(t, tt.args.name, got.GetName())
			assert.Equal(t, tt.text, got.String())
		})
	}
}
