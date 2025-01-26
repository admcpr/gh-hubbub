package auth

import (
	"fmt"
	"gh-reponark/shared"
	user1 "gh-reponark/user"
	"os/user"

	"github.com/charmbracelet/bubbles/v2/spinner"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/cli/go-gh/v2/pkg/api"
)

type AuthenticationErrorMsg struct{ Err error }
type AuthenticatedMsg struct{ User user.User }

type Model struct {
	spinner spinner.Model
	error   error
	// width   int
	// height  int
}

func (m *Model) SetDimensions(width, height int) {
	// m.width = width
	// m.height = height
}

func NewModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return Model{
		spinner: s,
	}
}

func (m Model) Init() (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{m.spinner.Tick, getUser}

	return m, tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case AuthenticationErrorMsg:
		m.error = msg.Err
	}

	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m Model) View() string {
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
	response := user1.User{}

	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return AuthenticationErrorMsg{Err: err}
	}

	return shared.NextMessage{ModelData: response}
}
