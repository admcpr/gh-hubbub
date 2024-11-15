package models

import (
	"gh-hubbub/structs"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
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
	// stack.Push(NewFiltersModel())

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

func (m MainModel) Init() tea.Cmd {
	child, _ := m.stack.Peek()
	return child.Init()
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

func (m MainModel) NextModel(message NextMessage) tea.Model {
	var newModel tea.Model
	switch m.stack.TypeOfHead() {
	case reflect.TypeOf(AuthenticatingModel{}):
		newModel = NewUserModel(message.ModelData.(structs.User), m.width, m.height)
	case reflect.TypeOf(UserModel{}):
		newModel = NewOrgModel(message.ModelData.(string), m.width, m.height)
	case reflect.TypeOf(OrgModel{}):
		newModel = NewFiltersModel()
	}

	return newModel
}
