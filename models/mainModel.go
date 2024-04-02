package models

import (
	"fmt"
	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh/v2/pkg/api"
)

var MainModel []tea.Model

type State int

const (
	Authenticating State = iota
	GettingOrgs
	ListingOrgs
	GettingRepos
	ListingRepos
	FilteringRepos
	EditingRepoFilter
)

func (s State) Next() {
	if s != EditingRepoFilter {
		s++
	}
}

func (s State) Previous() {
	if s != Authenticating {
		s--
	}
}

func (s State) Update(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.Previous()
		case "enter":
			s.Next()
		}
	}
}

type AuthenticationErrorMsg struct{ Err error }
type AuthenticatedMsg struct{ User structs.User }

type MainModelV2 struct {
	spinner   spinner.Model
	UserModel UserModel
	OrgModel  OrgModel
	RepoModel RepoModel

	state State
}

func NewMainModelV2() MainModelV2 {

	for i := 0; i < 10; i++ {

	}

	s := spinner.New()
	s.Spinner = spinner.Dot
	return MainModelV2{
		spinner: s,
		state:   Authenticating,
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
	case AuthenticatedMsg:
		m.UserModel = NewUserModel(msg.User)
		m.state = GettingOrgs
		return m.UserModel, m.UserModel.Init()

	default:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m MainModelV2) View() string {
	switch m.state {
	case Authenticating:
		return fmt.Sprintf("%s Authenticating ... \n\n", m.spinner.View())
	case GettingOrgs, ListingOrgs:
		return m.UserModel.View()
	case GettingRepos, ListingRepos, FilteringRepos, EditingRepoFilter:
		return m.RepoModel.View()
	default:
		return "Unknown state"
	}
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

	return AuthenticatedMsg{User: response}
}
