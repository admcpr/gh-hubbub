package models

import (
	"fmt"
	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh/v2/pkg/api"
)

var MainModel []tea.Model

type MainModelV2 struct {
	spinner   spinner.Model
	UserModel UserModel
	OrgModel  OrgModel
	RepoModel RepoModel

	User structs.User
}

func NewMainModelV2() MainModelV2 {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return MainModelV2{
		spinner: s,
	}
}

func (m MainModelV2) Init() tea.Cmd {
	cmds := []tea.Cmd{m.spinner.Tick, getUser}

	return tea.Batch(cmds...)
}

func (m MainModelV2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}
	default:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m MainModelV2) View() string {
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

	return AuthenticationMsg{User: response}
}
