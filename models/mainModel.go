package models

import (
	"gh-hubbub/structs"
	"reflect"
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type NextMessage struct{ ModelData interface{} }
type PreviousMessage struct{ ModelData interface{} }

type MainModel struct {
	stack  ModelStack
	width  int
	height int
}

func NewMainModel() MainModel {
	stack := ModelStack{}
	// stack.Push(NewAuthenticatingModel())
	// stack.Push(NewBoolModel("Is something true", false, 0, 0))
	stack.Push(NewDateModel("Date between", time.Now(), time.Now().Add(time.Hour*24*7), 0, 0))
	// stack.Push(NewIntModel("Number between", 0, 100, 0, 0))

	return MainModel{
		stack: stack,
	}
}

func (m *MainModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
	m.stack.SetDimensions(width, height)
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
		m.SetDimensions(msg.Width, msg.Height)
		return m, nil

	case tea.KeyPressMsg:
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
		cmd = m.Previous(msg)
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
	head, _ := m.stack.Peek()

	switch head.(type) {
	case AuthenticatingModel:
		newModel = NewUserModel(message.ModelData.(structs.User), m.width, m.height)
	case UserModel:
		newModel = NewOrgModel(message.ModelData.(string), m.width, m.height)
	case OrgModel:
		newModel = NewFiltersModel(m.width, m.height)
	}

	newModel, cmd := newModel.Init()
	m.stack.Push(newModel)

	return cmd
}

func (m *MainModel) Previous(message PreviousMessage) tea.Cmd {
	head, err := m.stack.Pop()

	if err != nil {
		return tea.Quit
	}

	switch head.(type) {
	case FiltersModel:
		// This is all a big mess, need to refactor to something less stinky
		if message.ModelData != nil && reflect.TypeOf(message.ModelData) == reflect.TypeOf(filterMap{}) {
			return m.UpdateChild(filtersMsg(message.ModelData.(filterMap)))
		}
	}

	return nil
}
