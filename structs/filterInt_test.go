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
		{args: args{tab: "tab", name: "name", from: 0, to: 0}, text: "name between 0 and 0"},
		{args: args{tab: "tab", name: "name", from: 0, to: 1}, text: "name between 0 and 1"},
		{args: args{tab: "tab", name: "name", from: 1, to: 0}, text: "name between 1 and 0"},
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

func TestFilterInt_Matches(t *testing.T) {
	zero, ten, hundred := 0, 10, 100

	tests := []struct {
		name    string
		filter  FilterInt
		setting Setting
		want    bool
	}{
		{
			name:    "value between from and to returns true",
			filter:  NewFilterInt("", "", zero, hundred),
			setting: NewSetting("", ten),
			want:    true,
		},
		{
			name:    "value equal from returns true",
			filter:  NewFilterInt("", "", zero, hundred),
			setting: NewSetting("", zero),
			want:    true,
		},
		{
			name:    "value equal to returns true",
			filter:  NewFilterInt("", "", zero, hundred),
			setting: NewSetting("", hundred),
			want:    true,
		},
		{
			name:    "value less than from returns false",
			filter:  NewFilterInt("", "", ten, hundred),
			setting: NewSetting("", zero),
			want:    false,
		},
		{
			name:    "value greater than to returns false",
			filter:  NewFilterInt("", "", zero, ten),
			setting: NewSetting("", hundred),
			want:    false,
		},
		{
			name:    "wrong type returns false",
			filter:  NewFilterInt("", "", zero, hundred),
			setting: NewSetting("", "it's a string"),
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.filter.Matches(tt.setting)
			assert.Equal(t, tt.want, got)
		})
	}
}
