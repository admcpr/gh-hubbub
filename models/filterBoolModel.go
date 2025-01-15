package models

import (
	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type BoolModel struct {
	Name   string
	Value  bool
	width  int
	height int
}

func NewBoolModel(name string, value bool, width, height int) BoolModel {
	m := BoolModel{
		Name:  name,
		Value: value,
	}

	m.width = width
	m.height = height

	return m
}

type BoolFilterMessage struct {
	Name  string
	Value bool
}

func (m *BoolModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
}

func (m BoolModel) Init() (tea.Model, tea.Cmd) {
	return m, textinput.Blink
}

func (m BoolModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			return m, m.SendFilterMsg
		case "esc":
			return m, func() tea.Msg {
				return PreviousMessage{}
			}
		case "y", "Y":
			m.Value = true
		case "n", "N":
			m.Value = false
		case "right", "left":
			m.Value = !m.Value
		}
	}

	return m, cmd
}

func (m BoolModel) View() string {
	yesButtonStyle := buttonStyle
	noButtonStyle := buttonStyle
	if m.Value {
		yesButtonStyle = activeButtonStyle
	} else {
		noButtonStyle = activeButtonStyle
	}
	buttons := lipgloss.JoinHorizontal(lipgloss.Left, yesButtonStyle.Render("Yes"), noButtonStyle.Render("No"))
	contents := lipgloss.JoinVertical(lipgloss.Center, modalTitleStyle.Render(m.Name), buttons)

	return lipgloss.PlaceHorizontal(m.width, lipgloss.Center, modalStyle.Render(contents))
}

func (m *BoolModel) GetValue() bool {
	return m.Value
}

func (m BoolModel) SendFilterMsg() tea.Msg {
	return PreviousMessage{ModelData: structs.NewFilterBool(m.Name, m.GetValue())}
}
