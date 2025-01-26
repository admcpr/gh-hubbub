package main

import (
	"gh-reponark/auth"
	"gh-reponark/filters"
	"gh-reponark/org"
	"gh-reponark/shared"
	"gh-reponark/user"

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
	stack.Push(auth.NewModel())
	// stack.Push(NewBoolModel("Is something true", false, 0, 0))
	// stack.Push(NewDateModel("Date between", time.Now(), time.Now().Add(time.Hour*24*7), 0, 0))
	// stack.Push(NewIntModel("Number between", 0, 100, 0, 0))
	// stack.Push(filters.NewModel(0, 0))

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
	case auth.Model:
		newModel = user.NewModel(message.ModelData.(user.User), m.width-2, m.height-2)
	case user.Model:
		newModel = org.NewModel(message.ModelData.(string), m.width-2, m.height-2)
	case org.Model:
		newModel = filters.NewModel(m.width-2, m.height-2)
	case filters.Model:
		newModel = filters.NewFilterModel(message.ModelData.(filters.Property), m.width-2, m.height-2)
	}

	newModel, cmd := newModel.Init()
	m.stack.Push(newModel)

	return cmd
}

func (m *MainModel) Previous(message shared.PreviousMessage) tea.Cmd {
	_, err := m.stack.Pop()

	if err != nil {
		return tea.Quit
	}

	if message.Message != nil {
		return m.UpdateChild(message.Message)
	}

	return nil
}
