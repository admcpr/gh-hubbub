package filters

import (
	"gh-hubbub/structs"
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
		{args: args{tab: "tab", name: "name", from: "2000-01-01", to: "2000-01-02"}, text: "name between 2000-01-01 and 2000-01-02"},
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

func TestFilterDate_Matches(t *testing.T) {
	now, yesterday, tomorrow := time.Now(), time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 1)

	tests := []struct {
		name    string
		filter  FilterDate
		setting structs.Setting
		want    bool
	}{
		{
			name:    "value between from and to returns true",
			filter:  NewFilterDate("", "", yesterday, tomorrow),
			setting: structs.NewSetting("", now),
			want:    true,
		},
		{
			name:    "value equal from returns true",
			filter:  NewFilterDate("", "", yesterday, tomorrow),
			setting: structs.NewSetting("", yesterday),
			want:    true,
		},
		{
			name:    "value equal to returns true",
			filter:  NewFilterDate("", "", yesterday, tomorrow),
			setting: structs.NewSetting("", tomorrow),
			want:    true,
		},
		{
			name:    "value less than from returns false",
			filter:  NewFilterDate("", "", now, tomorrow),
			setting: structs.NewSetting("", yesterday),
			want:    false,
		},
		{
			name:    "value greater than to returns false",
			filter:  NewFilterDate("", "", yesterday, now),
			setting: structs.NewSetting("", tomorrow),
			want:    false,
		},
		{
			name:    "wrong type returns false",
			filter:  NewFilterDate("", "", yesterday, tomorrow),
			setting: structs.NewSetting("", "it's a string"),
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
