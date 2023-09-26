package models

import (
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestOrgModel_Update_Quit(t *testing.T) {
	defaultOrgModel := NewOrgModel("admcpr", 100, 100)

	type args struct {
		msg tea.Msg
	}
	tests := []struct {
		name    string
		m       OrgModel
		args    args
		wantCmd tea.Cmd
	}{
		// TODO: Add more test cases.
		{"Quit with 'ctrl+c'", defaultOrgModel, args{tea.KeyMsg{Type: tea.KeyCtrlC}}, tea.Quit},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotCmd := tt.m.Update(tt.args.msg)
			if reflect.ValueOf(gotCmd) != reflect.ValueOf(tt.wantCmd) {
				t.Errorf("OrganisationModel.Update() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}
		})
	}
}

func TestOrgModel_NewOrgModel(t *testing.T) {
	type args struct {
		title  string
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test NewOrgModel", args: args{title: "gh-hubbub", width: 1024, height: 768}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOrgModel(tt.args.title, tt.args.width, tt.args.height)
			if got.Title != tt.args.title {
				t.Errorf("NewOrgModel title got = %v, want %v", got.Title, tt.args.title)
			}
			if got.width != tt.args.width {
				t.Errorf("NewOrgModel width got = %v, want %v", got.width, tt.args.width)
			}
			if got.height != tt.args.height {
				t.Errorf("NewOrgModel height got = %v, want %v", got.height, tt.args.height)
			}
			if got.keys.Enter.Enabled() != true {
				t.Errorf("NewOrgModel keys.Enter got = %v, want %v", got.keys.Enter.Enabled(), true)
			}
		})
	}
}

// func TestOrgModel_FilteredRepositories(t *testing.T) {
// 	repositories := []structs.Repository{
// 		structs.Repository{ Name: "name", Url: "url", SettingsTabs: []structs.SettingsTab{ structs.SettingsTab{ Name: "name", Settings: []structs.Setting{
// 			structs.Setting{ Name: "name", Value: "value", Type: "type", Description: "description", Url: "url" },
// 		} } } },
// 	OrgModel := NewOrgModel("admcpr", 100, 100)

// 	tests := []struct {
// 		name string
// 		m    *OrgModel
// 		want []structs.Repository
// 	}{
// 		"No filter"
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.m.FilteredRepositories(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("OrgModel.FilteredRepositories() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
