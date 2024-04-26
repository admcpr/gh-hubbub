package models

import (
	"fmt"
	"gh-hubbub/consts"
	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh/v2/pkg/api"
)

var MainModel []tea.Model

type State int

const (
	Authenticating State = iota
	ListingOrgs
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

type AuthenticationErrorMsg struct{ Err error }
type AuthenticatedMsg struct{ User structs.User }

type MainModelV2 struct {
	spinner   spinner.Model
	UserModel UserModel
	OrgModel  OrgModel
	RepoModel RepoModel

	state State

	width  int
	height int
}

func NewMainModelV2() MainModelV2 {
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

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.state == ListingOrgs {
				return m, tea.Quit
			}
			m.state.Previous()
		case "enter":
			switch m.state {
			case ListingOrgs:
				selectedName := m.UserModel.SelectedOrg().Login
				m.OrgModel = NewOrgModel(selectedName, m.width, m.height)
				m.state = ListingRepos
				cmd = m.OrgModel.Init()
				return m, cmd
			case ListingRepos:
				m.OrgModel.focus = consts.FocusTabs
			}
		case "ctrl+c":
			return m, tea.Quit
			// default:
			// 	switch m.state {
			// 	case ListingOrgs:
			// 		m.UserModel, cmd = m.UserModel.Update(msg)
			// 	}
		}
	case AuthenticatedMsg:
		m.UserModel = NewUserModel(msg.User, m.width, m.height)
		m.state = ListingOrgs
		return m, m.UserModel.Init()

	}

	switch m.state {
	case Authenticating:
		m.spinner, cmd = m.spinner.Update(msg)
	case ListingOrgs:
		m.UserModel, cmd = m.UserModel.Update(msg)
	case ListingRepos:
		m.OrgModel, cmd = m.OrgModel.Update(msg)
	}

	return m, cmd
}

func (m MainModelV2) View() string {
	switch m.state {
	case Authenticating:
		return fmt.Sprintf("%s Authenticating ... \n\n", m.spinner.View())
	case ListingOrgs:
		return m.UserModel.View()
	case ListingRepos, FilteringRepos, EditingRepoFilter:
		return m.OrgModel.View()
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
