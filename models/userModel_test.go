package models

import (
	"reflect"
	"testing"

	messages "gh-hubbub/messages"
	structs "gh-hubbub/structs"

	tea "github.com/charmbracelet/bubbletea"
)

func TestUserModel_Update(t *testing.T) {
	testUser := structs.User{Login: "test"}
	type args struct {
		msg tea.Msg
	}
	tests := []struct {
		name      string
		m         UserModel
		args      args
		wantModel tea.Model
		wantCmd   tea.Cmd
	}{
		// TODO: Add test cases.
		{"Authentication Success", UserModel{}, args{messages.AuthenticationMsg{User: testUser}}, UserModel{Authenticated: true, User: testUser}, getOrganisations},
		{"Authentication Failure", UserModel{}, args{messages.AuthenticationErrorMsg{}}, UserModel{Authenticated: false}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotModel, gotCmd := tt.m.Update(tt.args.msg)
			if !reflect.DeepEqual(gotModel, tt.wantModel) {
				t.Errorf("UserModel.Update() gotModel = %v, want %v", gotModel, tt.wantModel)
			}
			if reflect.ValueOf(gotCmd) != reflect.ValueOf(tt.wantCmd) {
				t.Errorf("UserModel.Update() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}
		})
	}
}
