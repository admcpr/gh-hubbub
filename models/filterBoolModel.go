package models

import (
	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type BoolModel struct {
	Name  string
	input textinput.Model
}

func NewBoolModel(name string, value bool) BoolModel {
	m := BoolModel{
		Name:  name,
		input: textinput.New(),
	}

	m.input.SetValue(structs.YesNo(value))
	m.input.Focus()

	return m
}

type BoolFilterMessage struct {
	Name  string
	Value bool
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
		case "y", "Y":
			m.input.SetValue("Yes")
		case "n", "N":
			m.input.SetValue("No")
		}
	}

	return m, cmd
}

func (m BoolModel) View() string {
	return m.Name + " " + m.input.View()
}

func (m *BoolModel) GetValue() bool {
	return m.input.Value() == "Yes"
}

func (m *BoolModel) Focus() tea.Cmd {
	return m.input.Focus()
}

func (m BoolModel) SendFilterMsg() tea.Msg {
	return structs.NewFilterBool(m.Name, m.GetValue())
}
