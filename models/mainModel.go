package models

import (
	"gh-hubbub/structs"

	tea "github.com/charmbracelet/bubbletea"
)

var MainModel []tea.Model

type State int

const (
	Authenticating State = iota
	ListingOrgs
	ListingRepos
	SelectedRepo
	FilteringRepos
	EditingRepoFilter
)

func (s State) Next() State {
	if s != EditingRepoFilter {
		return s + 1
	}
	return s
}

func (s State) Previous() State {
	if s != Authenticating {
		return s - 1
	}
	return s
}

type AuthenticationErrorMsg struct{ Err error }
type AuthenticatedMsg struct{ User structs.User }
type NextMessage struct{ ModelData interface{} }
type PreviousMessage struct{}

type MainModelV2 struct {
	stack     Stack
	models    []tea.Model
	authModel AuthenticatingModel
	UserModel UserModel
	OrgModel  OrgModel
	RepoModel RepoModel
	yes       bool
	width     int
	height    int
}

func NewMainModelV2() MainModelV2 {
	stack := Stack{}
	stack.Push(NewAuthenticatingModel())

	return MainModelV2{
		stack: stack,
	}
}

func (m *MainModelV2) SetWidth(width int) {
	m.width = width
	m.UserModel.SetWidth(width)
	m.OrgModel.SetWidth(width)
	m.RepoModel.SetWidth(width)
}

func (m *MainModelV2) SetHeight(height int) {
	m.height = height
	m.UserModel.SetHeight(height)
	m.OrgModel.SetHeight(height)
	m.RepoModel.SetHeight(height)
}

func (m MainModelV2) Init() tea.Cmd {
	return m.authModel.Init()
}

func (m MainModelV2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// This should only do a couple of things
	// 1. Handle ctrl+c to quit
	// 2. Handle window sizing
	// 3. Handle Forward & Back navigation (creating models as needed) and updating state
	// 4. Call Update on the active model

	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.SetHeight(msg.Height)
		m.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		// case "esc":
		// 	if m.state == ListingOrgs {
		// 		return m, tea.Quit
		// 	}
		// 	m.state = m.state.Previous()
		// 	return m, nil
		// case "enter":
		// 	switch m.state {
		// 	case ListingOrgs:
		// 		selectedName := m.UserModel.SelectedOrg().Login
		// 		m.OrgModel = NewOrgModel(selectedName, m.width, m.height)
		// 		m.state = ListingRepos
		// 		cmd = m.OrgModel.Init()
		// 		return m, cmd
		// 	case ListingRepos:
		// 		m.state = SelectedRepo
		// 		m.OrgModel.focus = consts.FocusTabs
		// 		return m, cmd
		// 	}
		case "ctrl+c":
			return m, tea.Quit
		}
	case NextMessage:
		m.yes = true
		return m, nil
	}

	// switch m.state {
	// case ListingOrgs:
	// 	m.UserModel, cmd = m.UserModel.Update(msg)
	// case ListingRepos:
	// 	m.OrgModel, cmd = m.OrgModel.Update(msg)
	// }
	m.authModel, cmd = m.authModel.Update(msg)

	return m, cmd
}

func (m MainModelV2) View() string {
	if m.yes {
		return "Yes"
	}
	return m.authModel.View()
	// switch m.state {
	// case ListingOrgs:
	// 	return m.UserModel.View()
	// case ListingRepos, FilteringRepos, EditingRepoFilter, SelectedRepo:
	// 	return m.OrgModel.View()
	// default:
	// 	return "Unknown state"
	// }
}
