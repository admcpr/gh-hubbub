package models

import (
	"testing"

	structs "gh-hubbub/structs"
)

func TestTabs_RenderTabs(t *testing.T) {
	threeTabs := []structs.SettingsTab{
		{Name: "Tab 1"},
		{Name: "Tab 2"},
		{Name: "Tab 3"},
	}

	type args struct {
		tabSettings []structs.SettingsTab
		width       int
		activeTab   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Three Tabs, Active 0",
			args: args{
				tabSettings: threeTabs,
				width:       35,
				activeTab:   0,
			},
			want: " ─────────────────────────────────── \n    Tab 1      Tab 2       Tab 3     "},
		{
			name: "Three Tabs, Active 1",
			args: args{
				tabSettings: threeTabs,
				width:       30,
				activeTab:   1,
			},
			want: " ─────────────────────────── \n   Tab 1    Tab 2    Tab 3   "},
		{
			name: "Three Tabs, Active 2",
			args: args{
				tabSettings: threeTabs,
				width:       25,
				activeTab:   2,
			},
			want: " ────────────────────── \n  Tab 1  Tab 2  Tab 3   "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RenderTabs(tt.args.tabSettings, tt.args.width, tt.args.activeTab); got != tt.want {
				t.Errorf("RenderTabs() = %v, want %v", got, tt.want)
			}
		})
	}
}
