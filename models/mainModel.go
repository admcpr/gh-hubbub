package models

import (
	"gh-hubbub/structs"
	"reflect"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type NextMessage struct{ ModelData interface{} }
type PreviousMessage struct{}

type MainModel struct {
	stack  Stack
	width  int
	height int
}

func NewMainModel() MainModel {
	stack := Stack{}
	stack.Push(NewAuthenticatingModel())

	return MainModel{
		stack: stack,
	}
}

func (m *MainModel) SetWidth(width int) {
	m.width = width
	// TODO: Set width of stack head
}

func (m *MainModel) SetHeight(height int) {
	m.height = height
	// TODO: Set height of stack head
}

func (m MainModel) Init() (tea.Model, tea.Cmd) {
	child, _ := m.stack.Peek()
	_, cmd := child.Init()
	return m, cmd
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// This should only do a couple of things
	// 1. Handle ctrl+c to quit ✔️
	// 2. Handle window sizing ✔️
	// 3. Handle Forward & Back navigation (creating models as needed) and updating state ✔️
	// 4. Call Update on the active model ✔️

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
		cmd = m.Next(msg)
		return m, cmd
	case PreviousMessage:
		cmd = m.Previous()
		return m, cmd
	default:
		cmd = m.UpdateChild(msg)
	}

	return m, cmd
}

func (m *MainModel) UpdateChild(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	currentModel, _ := m.stack.Pop()
	currentModel, cmd = currentModel.Update(msg)
	m.stack.Push(currentModel)
	return cmd
}

func (m MainModel) View() string {
	child, _ := m.stack.Peek()
	return child.View()
}

func (m *MainModel) Next(message NextMessage) tea.Cmd {
	var newModel tea.Model
	switch m.stack.TypeOfHead() {
	case reflect.TypeOf(AuthenticatingModel{}):
		newModel = NewUserModel(message.ModelData.(structs.User), m.width, m.height)
	case reflect.TypeOf(UserModel{}):
		newModel = NewOrgModel(message.ModelData.(string), m.width, m.height)
	case reflect.TypeOf(OrgModel{}):
		newModel = NewFiltersModel()
	}

	nextModel, cmd := newModel.Init()
	m.stack.Push(nextModel)

	return cmd
}

func (m *MainModel) Previous() tea.Cmd {
	_, err := m.stack.Pop()
	if err != nil {
		return tea.Quit
	}
	return nil
}
