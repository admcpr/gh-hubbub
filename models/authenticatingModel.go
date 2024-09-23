package models

import (
	"fmt"
	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh/v2/pkg/api"
)

type AuthenticationErrorMsg struct{ Err error }
type AuthenticatedMsg struct{ User structs.User }

type AuthenticatingModel struct {
	spinner spinner.Model
	error   error
}

func NewAuthenticatingModel() AuthenticatingModel {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return AuthenticatingModel{
		spinner: s,
	}
}

func (m AuthenticatingModel) Init() tea.Cmd {
	cmds := []tea.Cmd{m.spinner.Tick, getUser}

	return tea.Batch(cmds...)
}

func (m AuthenticatingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case AuthenticationErrorMsg:
		m.error = msg.Err
	}

	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m AuthenticatingModel) View() string {
	if m.error != nil {
		return fmt.Sprintf("Error authenticating: %s\nctrl+c to exit", m.error)
	}
	return fmt.Sprintf("%s Authenticating ... \n\n", m.spinner.View())
}

func getUser() tea.Msg {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return AuthenticationErrorMsg{Err: err}
	}
	response := structs.User{}

	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return AuthenticationErrorMsg{Err: err}
	}

	return NextMessage{ModelData: response}
}
