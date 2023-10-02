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
		{args: args{tab: "tab", name: "name", value: true}, text: "name = Yes"},
		{args: args{tab: "tab", name: "name", value: false}, text: "name = No"},
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

func TestFilterBool_Matches(t *testing.T) {
	tests := []struct {
		name    string
		filter  FilterBool
		setting Setting
		want    bool
	}{
		{
			name:    "matching value returns true",
			filter:  NewFilterBool("", "", true),
			setting: NewSetting("", true),
			want:    true,
		},
		{
			name:    "non-matching value returns false",
			filter:  NewFilterBool("", "", false),
			setting: NewSetting("", false),
			want:    true,
		},
		{
			name:    "wrong type returns false",
			filter:  NewFilterBool("", "", false),
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
