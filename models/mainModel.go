package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

var MainModel []tea.Model

type MainModelV2 struct {
	spinner   spinner.Model
	UserModel UserModel
	OrgModel  OrgModel
	RepoModel RepoModel
}

func NewMainModelV2() MainModelV2 {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return MainModelV2{
		spinner: s,
	}
}

func (m MainModelV2) Init() tea.Cmd {
	return m.spinner.Tick
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
