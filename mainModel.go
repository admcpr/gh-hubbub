package main

import (
	"gh-reponark/filters"
	"gh-reponark/orgs"
	"gh-reponark/shared"
	"gh-reponark/users"
	"reflect"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type MainModel struct {
	stack  shared.ModelStack
	width  int
	height int
}

func NewMainModel() MainModel {
	stack := shared.ModelStack{}
	// stack.Push(NewAuthenticatingModel())
	// stack.Push(NewBoolModel("Is something true", false, 0, 0))
	// stack.Push(NewDateModel("Date between", time.Now(), time.Now().Add(time.Hour*24*7), 0, 0))
	// stack.Push(NewIntModel("Number between", 0, 100, 0, 0))
	stack.Push(filters.NewModel(0, 0))

	return MainModel{
		stack: stack,
	}
}

func (m *MainModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
	// 2 is subtracted from the width and height to account for the border
	m.stack.SetDimensions(width-2, height-2)
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
	case shared.NextMessage:
		cmd = m.Next(msg)
		return m, cmd
	case shared.PreviousMessage:
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
	borderStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(shared.AppColors.Green)

	return borderStyle.Render(lipgloss.PlaceHorizontal(m.width-2, lipgloss.Left, child.View()))
}

func (m *MainModel) Next(message shared.NextMessage) tea.Cmd {
	var newModel tea.Model
	head, _ := m.stack.Peek()

	switch head.(type) {
	case users.AuthenticatingModel:
		newModel = users.NewUserModel(message.ModelData.(users.User), m.width-2, m.height-2)
	case users.UserModel:
		newModel = orgs.NewModel(message.ModelData.(string), m.width-2, m.height-2)
	case orgs.Model:
		newModel = filters.NewModel(m.width-2, m.height-2)
	case filters.Model:
		newModel = filters.NewFilterModel(message.ModelData.(filters.Property), m.width-2, m.height-2)
	}

	newModel, cmd := newModel.Init()
	m.stack.Push(newModel)

	return cmd
}

func (m *MainModel) Previous(message shared.PreviousMessage) tea.Cmd {
	head, err := m.stack.Pop()

	if err != nil {
		return tea.Quit
	}

	// TODO: see if this works instead of the madness below
	// if message.ModelData != nil {
	// 	return m.UpdateChild(message.ModelData)
	// }

	switch head.(type) {
	case filters.Model:
		// This is all a big mess, need to refactor to something less stinky
		if message.ModelData != nil && reflect.TypeOf(message.ModelData) == reflect.TypeOf(filters.FilterMap{}) {
			return m.UpdateChild(filters.FiltersMsg(message.ModelData.(filters.FilterMap)))
		}
	case filters.IntModel, filters.DateModel, filters.BoolModel:
		if message.ModelData != nil {
			return m.UpdateChild(filters.AddFilterMsg(message.ModelData.(filters.Filter)))
		}
	}

	return nil
}
