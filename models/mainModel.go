package models

import (
	"gh-hubbub/structs"
	"reflect"

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
	stack  Stack
	width  int
	height int
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
	// TODO: Set width of stack head
}

func (m *MainModelV2) SetHeight(height int) {
	m.height = height
	// TODO: Set height of stack head
}

func (m MainModelV2) Init() tea.Cmd {
	child, _ := m.stack.Peek()
	return child.Init()
}

func (m MainModelV2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// This should only do a couple of things
	// 1. Handle ctrl+c to quit
	// 2. Handle window sizing ✔️
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
		case "ctrl+c":
			return m, tea.Quit
		default:
			cmd = m.UpdateChild(msg)
		}
	case NextMessage:
		// TODO: This is where we should be pushing the new model onto the stack
		nextModel := m.NextModel(msg)
		m.stack.Push(nextModel)
		return m, nextModel.Init()
	case PreviousMessage:
		_, err := m.stack.Pop()
		if err != nil {
			return m, tea.Quit
		}
		return m, nil
	default:
		cmd = m.UpdateChild(msg)
	}

	return m, cmd
}

func (m *MainModelV2) UpdateChild(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	currentModel, _ := m.stack.Pop()
	currentModel, cmd = currentModel.Update(msg)
	m.stack.Push(currentModel)
	return cmd
}

func (m MainModelV2) View() string {
	child, _ := m.stack.Peek()
	return child.View()
}

func (m MainModelV2) NextModel(message NextMessage) tea.Model {
	var newModel tea.Model
	switch m.stack.TypeOfHead() {
	case reflect.TypeOf(AuthenticatingModel{}):
		newModel = NewUserModel(message.ModelData.(structs.User), m.width, m.height)
	case reflect.TypeOf(UserModel{}):
		newModel = NewOrgModel(message.ModelData.(string), m.width, m.height)
	}
	return newModel
}
